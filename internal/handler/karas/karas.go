package karas

import (
	"github.com/haytty/karas/cli/flags"
	"github.com/haytty/karas/internal/store"
	"github.com/haytty/karas/internal/webdriver"

	"github.com/haytty/karas/internal/model"
)

func Apply(option *flags.GlobalOption) error {
	w := webdriver.NewSelenium(option.SeleniumPath, option.Chrome, option.ChromeDriver, option.Port)
	m := model.NewKaras(option.Config, option.JSON, w)
	s := store.NewStore()

	if err := m.Load(); err != nil {
		return err
	}

	if err := m.Do(); err != nil {
		return err
	}
	s.Dump("")
	return nil
}
