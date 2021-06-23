package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

var usage = ` Usage:
	Dmp timely --config=<x.toml> [--test]
	Dmp daily  --config=<x.toml> [--test]
	Dmp sampling --config=<x.toml> [--test]
	Dmp -h | --help
	Dmp -v | --version

Options:
	timely                       Update the followlist cache in real time.
	daily                        Update the watchlist cache by day.
	sampling                     Sampling consistency check.
	-h --help                    Show this screen.
	-v --version                 Show version.
`

type cmd struct {
	IsCmdTimely   bool   `docopt:"timely"`
	IsCmdDaily    bool   `docopt:"daily"`
	IsCmdSampling bool   `docopt:"sampling"`
	CfgFname      string `docopt:"--config"`
	IsTest        bool   `docopt:"--test"`
	ISHelp        bool   `docopt:"--help"`
	ISVersion     bool   `docopt:"--version"`
}

func main() {
	fmt.Println("vim-go")
	var c cmd
	err := c.Run()
	fmt.Println("text from main.Run")
	fmt.Println(err)
}

func (c *cmd) Run() error {
	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Println("??")
		fmt.Println(err)
		fmt.Println(usage)
		return err
	}
	fmt.Println(opts)
	if err != nil {
		return err
	}
	if c.IsCmdTimely {
		return c.TimelyEntry()
	}
	return nil
}
func (c *cmd) TimelyEntry() error {
	if c.IsCmdTimely {
		fmt.Println("timely start")
	}
	return nil
}
