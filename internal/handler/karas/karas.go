package karas

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/karas/internal/model"
)

func Apply() error {
	karas := model.NewKaras()

	fmt.Printf(
		color.GreenString("Root Action:%s \n", karas.String()),
	)

	return nil
}
