package event

import (
	"fmt"

	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
	"github.com/tebeka/selenium"
)

type InputEvent struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}

func (i InputEvent) Act(wd selenium.WebDriver) error {
	elem, err := wd.FindElement(i.Selector.Type.TypeName(), i.Selector.Value.Value())
	if err != nil {
		return fmt.Errorf("input event: find element error: %v", err)
	}

	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		return fmt.Errorf("input event: clear text error: %v", err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys(i.Value.String())
	if err != nil {
		return fmt.Errorf("input event: send string error: %v", err)
	}

	return nil
}
