package event

import (
	"fmt"

	"github.com/haytty/karas/internal/model/action/page"

	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
	"github.com/tebeka/selenium"
)

type ClickEvent struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}

func (e ClickEvent) Act(wd selenium.WebDriver) error {
	beforePage, btn, err := e.collectPageInfo(wd)
	if err != nil {
		return err
	}

	if err := btn.Click(); err != nil {
		return fmt.Errorf("click event: button click error: %v", err)
	}

	if err := wd.WaitWithTimeoutAndInterval(e.check(beforePage), defaultWaitTime, defaultIntervalTime); err != nil {
		return err
	}

	return nil
}

func (e ClickEvent) collectPageInfo(wd selenium.WebDriver) (*page.Page, selenium.WebElement, error) {
	url, err := wd.CurrentURL()
	if err != nil {
		return nil, nil, fmt.Errorf("click event: retrieve url error: %v", err)
	}
	p := page.NewPage()
	p.URL = url

	elem, err := wd.FindElement(e.Selector.Type.TypeName(), e.Selector.Value.Value())
	if err != nil {
		return nil, nil, fmt.Errorf("click event: find element error: %v", err)
	}

	return p, elem, nil
}

func (e ClickEvent) check(beforePage *page.Page) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		afterPage, _, err := e.collectPageInfo(wd)
		if err != nil {
			return false, err
		}

		if !afterPage.Match(beforePage) {
			return false, fmt.Errorf("click event: find element error: %v", err)
		}

		return true, nil
	}
}
