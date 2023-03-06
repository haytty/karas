package event

import (
	"fmt"

	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
	"github.com/tebeka/selenium"
)

type SubmitEvent struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}

func (i SubmitEvent) Act(wd selenium.WebDriver) error {
	btn, err := wd.FindElement(i.Selector.Type.TypeName(), i.Selector.Value.Value())
	if err != nil {
		return fmt.Errorf("submit event: find element error: %v", err)
	}

	if err := btn.Submit(); err != nil {
		return fmt.Errorf("submit event: button submit error: %v", err)
	}

	return nil
}
