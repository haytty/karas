package event

import (
	"time"

	"github.com/tebeka/selenium"
)

const (
	defaultWaitTime     = 10 * time.Second
	defaultIntervalTime = 1 * time.Second
)

type Event interface {
	Act(selenium.WebDriver) error
}
