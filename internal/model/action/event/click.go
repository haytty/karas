package event

import (
	"fmt"

	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
	"github.com/tebeka/selenium"
)

type ClickEvent struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}

func (i ClickEvent) Act(wd selenium.WebDriver) error {
	btn, err := wd.FindElement(i.Selector.Type.TypeName(), i.Selector.Value.Value())
	if err != nil {
		return fmt.Errorf("click event: find element error: %v", err)
	}

	if err := btn.Click(); err != nil {
		return fmt.Errorf("click event: button click error: %v", err)
	}

	return nil
}
