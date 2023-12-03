package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

const pathListSeparator = string(os.PathSeparator)

type Entry interface {
	// ReadClass /** 负责寻找和加载class文件*/
	ReadClass(className string) ([]byte, Entry, error)
	// String /**返回变量的字符串表示*/
	String() string
}

func NewEntry(path string) Entry {
	//if strings.Contains(path, pathListSeparator) {
	//	return newCompositeEntry(path)
	//}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, "*.jar") || strings.HasSuffix(path, "*.zip") ||
		strings.HasSuffix(path, "*.ZIP") || strings.HasSuffix(path, "*.JAR") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

type Classpath struct {
	BootClassPath Entry
	ExtClassPath  Entry
	UserClassPath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.ParseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.BootClassPath.ReadClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.ExtClassPath.ReadClass(className); err == nil {
		return data, entry, err
	}
	return self.UserClassPath.ReadClass(className)
}
func (self *Classpath) String() string {
	return self.UserClassPath.String()
}

func (self *Classpath) ParseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.BootClassPath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.ExtClassPath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("can't find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.UserClassPath = NewEntry(cpOption)
}
