package model

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/haytty/karas/internal/model/action"
	"github.com/tebeka/selenium"

	"github.com/haytty/karas/internal/webdriver"
)

type Karas struct {
	Suite     KarasSuite `json:"suite"`
	config    string
	json      string
	webdriver webdriver.WebDriver
}

type KarasSuite struct {
	Url     string          `json:"url"`
	Output  string          `json:"output"`
	Actions []action.Action `json:"actions"`
}

func NewKaras(config string, json string, driver webdriver.WebDriver) *Karas {
	return &Karas{
		config:    config,
		json:      json,
		webdriver: driver,
	}
}

func (k *Karas) Load() error {
	b, err := os.ReadFile(k.json)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, k); err != nil {
		return err
	}
	return nil
}

func (k *Karas) Do() error {
	service, wd, err := k.webdriver.NewWebDriver()
	if err != nil {
		return err
	}
	defer service.Stop()
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get(k.Suite.Url); err != nil {
		return err
	}

	// wait starting browser.
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.CurrentURL()
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return fmt.Errorf("browser open error: %v", err)
	}

	for _, action := range k.Suite.Actions {
		fmt.Printf("Execute Action: %s\n", action.Name)
		if err := action.Event.Act(wd); err != nil {
			return err
		}
	}
	return nil
}
