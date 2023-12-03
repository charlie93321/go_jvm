package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func (self DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data, self, err

}

func (self DirEntry) String() string {
	return self.absDir
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
