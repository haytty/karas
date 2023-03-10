package karas

import (
	"github.com/haytty/karas/cli/flags"
	"github.com/haytty/karas/internal/builder"
	"github.com/haytty/karas/internal/store"

	"github.com/haytty/karas/internal/model"
)

func Apply(option *flags.GlobalOption) error {
	var karas *model.Karas
	var err error

	if option.JSON != "" {
		karas, err = builder.BuildFromKarasJSON(option.JSON, option.SeleniumPath, option.Chrome, option.ChromeDriver, option.Port)
	} else {
		karas, err = builder.BuildFromConfig(option.Config, option.SeleniumPath, option.Chrome, option.ChromeDriver, option.Port)
	}

	s := store.NewStore()

	if err := karas.Do(); err != nil {
		return err
	}

	s.Dump(option.Format)

	return err
}
