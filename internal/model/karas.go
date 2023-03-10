package model

import (
	"fmt"
	"time"

	"github.com/haytty/karas/internal/model/action"
	"github.com/tebeka/selenium"

	"github.com/haytty/karas/internal/webdriver"
)

type Karas struct {
	Suite     KarasSuite `json:"suite"`
	webdriver webdriver.WebDriver
}

type KarasSuite struct {
	URL string `json:"url"`
	//ImplicitWaitTime time.Duration   `json:"implicitWaitTime"`
	Output  string          `json:"output"`
	Actions []action.Action `json:"actions"`
}

func NewKaras(driver webdriver.WebDriver) *Karas {
	return &Karas{
		webdriver: driver,
	}
}

const (
	defaultImplicitWaitTime = 5 * time.Second
)

func (k *Karas) Do() error {
	service, wd, err := k.webdriver.NewWebDriver()
	if err != nil {
		return err
	}
	defer service.Stop()
	defer wd.Quit()

	if err := wd.SetImplicitWaitTimeout(defaultImplicitWaitTime); err != nil {
		return err
	}

	// Navigate to the simple playground interface.
	if err := wd.Get(k.Suite.URL); err != nil {
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
