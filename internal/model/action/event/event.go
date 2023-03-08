package event

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

const (
	defaultWaitTime     = 30 * time.Second
	defaultIntervalTime = 1 * time.Second
)

type Event interface {
	Act(selenium.WebDriver) error
}

func findElement(wd selenium.WebDriver, selectorType, selectorValue string) (selenium.WebElement, error) {
	var elem selenium.WebElement
	var err error

	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		elem, err = wd.FindElement(selectorType, selectorValue)
		if err != nil {
			fmt.Printf("%v\n", err)
			return false, nil
		}
		return true, nil
	})
	return elem, err
}
