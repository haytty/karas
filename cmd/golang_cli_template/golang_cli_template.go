package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/haytty/golang_cli_template/cli"
	clistruct "github.com/haytty/golang_cli_template/cli/cli"
)

func main() {
	c := clistruct.NewGolangCliTemplateCli()
	if err := cli.NewGolangCliTemplateCommand(c).Execute(); err != nil {
		fmt.Fprintln(c.Out(), color.RedString(err.Error()))
		os.Exit(1)
	}
}
