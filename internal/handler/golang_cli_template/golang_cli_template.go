package golang_cli_template

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/golang_cli_template/internal/model"
)

func Apply() error {
	golang_cli_template := model.NewGolangCliTemplate()

	fmt.Printf(
		color.GreenString("Root Action:%s \n", golang_cli_template.String()),
	)

	return nil
}
