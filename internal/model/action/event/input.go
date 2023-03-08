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
	_, elem, err := e.collectPageInfo(wd)
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

	return nil
}

func (e InputEvent) collectPageInfo(wd selenium.WebDriver) (*page.Page, selenium.WebElement, error) {
	p := page.NewPage()

	elem, err := findElement(wd, e.Selector.Type.TypeName(), e.Selector.Value.Value())
	if err != nil {
		return nil, nil, fmt.Errorf("input event: find element error: %v", err)
	}

	return p, elem, nil
}
