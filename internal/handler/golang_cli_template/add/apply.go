package add

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/golang_cli_template/internal/model"
)

func Apply(param1, param2 string) error {
	golang_cli_template := model.NewGolangCliTemplate()

	fmt.Printf(
		color.GreenString("Add Action:%s \n") +
		color.GreenString("Arg1:%s \n") +
		color.GreenString("Arg2:%s \n"),
		golang_cli_template.String(), 
    param1, 
    param2, 
	)

	return nil
}
