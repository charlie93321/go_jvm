package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
}

func ParseCmd() *Cmd {
	cmd1 := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd1.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd1.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd1.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd1.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd1.CpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd1.Class = args[0]
		cmd1.Args = args[1:]
	}
	return cmd1
}

func PrintUsage() {
	fmt.Printf("Usage:.%s [-options] class [args...] \n", os.Args[0])
}

func StartJVM(c *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", c.CpOption, c.Class, c.Args)
}
