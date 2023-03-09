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

const (
	defaultRefreshCount = 2
)

func findElement(wd selenium.WebDriver, selectorType, selectorValue string) (selenium.WebElement, error) {
	return findElementWithRefresh(wd, selectorType, selectorValue, defaultRefreshCount)
}

func findElementWithRefresh(wd selenium.WebDriver, selectorType, selectorValue string, refreshCount int) (selenium.WebElement, error) {
	var elem selenium.WebElement
	var err error

	err = wd.WaitWithTimeoutAndInterval(func(wd selenium.WebDriver) (bool, error) {
		elem, err = wd.FindElement(selectorType, selectorValue)
		if err != nil {
			fmt.Printf("%v\n", err)
			return false, nil
		}
		return true, nil
	}, defaultWaitTime, defaultIntervalTime)
	if err != nil {
		refreshCount -= 1

		if refreshCount >= 0 {

			src, err := wd.PageSource()
			if err != nil {
				return elem, err
			}

			// DEBUG
			fmt.Println(src)

			// has no effect
			if err := wd.Refresh(); err != nil {
				return elem, err
			}

			fmt.Println("no such an element. page refresh!!!")

			return findElementWithRefresh(wd, selectorType, selectorValue, refreshCount)
		}
	}

	return elem, err
}
