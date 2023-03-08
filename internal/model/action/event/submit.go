package event

import (
	"fmt"

	"github.com/haytty/karas/internal/model/action/page"

	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
	"github.com/tebeka/selenium"
)

type SubmitEvent struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}

func (e SubmitEvent) Act(wd selenium.WebDriver) error {
	beforePage, btn, err := e.collectPageInfo(wd)
	if err != nil {
		return err
	}

	if err := btn.Submit(); err != nil {
		return fmt.Errorf("submit event: button submit error: %v", err)
	}

	if err := wd.WaitWithTimeoutAndInterval(e.check(beforePage), defaultWaitTime, defaultIntervalTime); err != nil {
		return err
	}

	return nil
}

func (e SubmitEvent) collectPageInfo(wd selenium.WebDriver) (*page.Page, selenium.WebElement, error) {
	url, err := wd.CurrentURL()
	if err != nil {
		return nil, nil, fmt.Errorf("submit event: retrieve url error: %v", err)
	}
	p := page.NewPage()
	p.URL = url

	elem, err := findElement(wd, e.Selector.Type.TypeName(), e.Selector.Value.Value())
	if err != nil {
		return nil, nil, fmt.Errorf("submit event: find element error: %v", err)
	}

	return p, elem, nil
}

func (e SubmitEvent) check(beforePage *page.Page) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		afterPage, _, err := e.collectPageInfo(wd)
		if err != nil {
			// Returns nil because polling is not performed when error contents are returned
			return false, err
		}

		if !afterPage.Match(beforePage) {
			// Returns nil because polling is not performed when error contents are returned
			return false, nil
		}

		return true, nil
	}
}
