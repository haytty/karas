package event

import "github.com/tebeka/selenium"

type Event interface {
	Act(selenium.WebDriver) error
}
