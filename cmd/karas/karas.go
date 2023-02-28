package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/haytty/karas/cli"
	clistruct "github.com/haytty/karas/cli/cli"
)

func main() {
	c := clistruct.NewKarasCli()
	if err := cli.NewKarasCommand(c).Execute(); err != nil {
		fmt.Fprintln(c.Out(), color.RedString(err.Error()))
		os.Exit(1)
	}
}
