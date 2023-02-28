package add

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/karas/internal/model"
)

func Apply(param1, param2 string) error {
	karas := model.NewKaras()

	fmt.Printf(
		color.GreenString("Add Action:%s \n")+
			color.GreenString("Arg1:%s \n")+
			color.GreenString("Arg2:%s \n"),
		karas.String(),
		param1,
		param2,
	)

	return nil
}
