package event

import (
	"fmt"

	"github.com/haytty/karas/internal/model/action/page"

	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
	"github.com/tebeka/selenium"
)

type InputEvent struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}

func (e InputEvent) Act(wd selenium.WebDriver) error {
	beforePage, elem, err := e.collectPageInfo(wd)
	if err != nil {
		return err
	}

	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		return fmt.Errorf("input event: clear text error: %v", err)
	}

	err = elem.SendKeys(e.Value.String())
	if err != nil {
		return fmt.Errorf("input event: send string error: %v", err)
	}

	if err := wd.WaitWithTimeoutAndInterval(e.check(beforePage), defaultWaitTime, defaultIntervalTime); err != nil {
		return err
	}

	return nil
}

func (e InputEvent) collectPageInfo(wd selenium.WebDriver) (*page.Page, selenium.WebElement, error) {
	p := page.NewPage()

	elem, err := wd.FindElement(e.Selector.Type.TypeName(), e.Selector.Value.Value())
	if err != nil {
		return nil, nil, fmt.Errorf("input event: find element error: %v", err)
	}

	t, err := elem.Text()
	if err != nil {
		return nil, nil, fmt.Errorf("input event: retrieve text error: %v", err)
	}

	p.Set(e.Selector.Value.Value(), t)

	return p, elem, nil
}

func (e InputEvent) check(beforePage *page.Page) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		afterPage, _, err := e.collectPageInfo(wd)
		if !afterPage.Match(beforePage) {
			return false, fmt.Errorf("input event: check element error: input value difference: %v", err)
		}

		return true, nil
	}
}
