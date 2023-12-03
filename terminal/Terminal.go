package terminal

import (
	"fmt"
	"go_jvm/classpath"
	"os"
	"strings"
)

type Terminal struct {
	HelpFlag    bool     `json:"help"`
	VersionFlag bool     `json:"version"`
	CpOption    string   `json:"classpath"`
	Class       string   `json:"class"`
	Args        []string `json:"args"`
	XJreOption  string   `json:"xjre"`
}

func ParseCmd(args []string) *Terminal {
	cmd1 := &Terminal{}
	args = removeSpace(args)
	if len(args) == 1 {
		if "--help" == args[0] || "?" == args[0] {
			cmd1.HelpFlag = true
		}

		if "--version" == args[0] || "-v" == args[0] {
			cmd1.VersionFlag = true
		}
	}
	if len(args) > 1 {
		cmd1.Class = args[0]
		cmd1.Args = args[1:]
		for index, value := range cmd1.Args {
			if value == "--classpath" || value == "--cp" {
				if len(cmd1.Args) > index+1 {
					cmd1.CpOption = cmd1.Args[index+1]
				}
			} else if value == "--xjre" {
				if len(cmd1.Args) > index+1 {
					cmd1.XJreOption = cmd1.Args[index+1]
				}
			}
		}
	}
	return cmd1
}

func removeSpace(args []string) []string {
	var arr []string
	for _, v := range args {
		is := strings.TrimSpace(v)
		if len(is) > 0 {
			arr = append(arr, is)
		}
	}
	return arr
}

func PrintUsage() {
	fmt.Printf("Usage:.%s [-options] class [args...] \n", os.Args[0])
}

func StartJVM(c *Terminal) {
	//fmt.Printf("classpath:%s class:%s args:%v\n", c.CpOption, c.Class, c.Args)

	cp := classpath.Parse(c.XJreOption, c.CpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", c.CpOption, c.Class, c.Args)
	className := strings.Replace(c.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("could't find or load main class %s\n", c.Class)
		return
	}
	fmt.Printf("class data :%v\n", string(classData))

}
