package main

import (
	"encoding/json"
	"fmt"
	"go_jvm/terminal"
)

func main() {
	println("............启动项目...........")

	args := []string{"com.dkd.shape.Hello",
		"--classpath", "D:\\app\\workspace\\stu-mysql\\out\\production\\stu-mysql",
		"", "--xjre", "D:\\app\\env\\jre8"}

	c := terminal.ParseCmd(args)
	jsonData, err := json.Marshal(c)
	fmt.Println(string(jsonData), err)
	if c.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if c.HelpFlag || c.Class == "" {
		terminal.PrintUsage()
	} else {
		terminal.StartJVM(c)
	}
}
