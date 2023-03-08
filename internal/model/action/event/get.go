package event

import (
	"fmt"

	"github.com/haytty/karas/internal/store"

	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
	"github.com/tebeka/selenium"
)

type GetEvent struct {
	Selector selector.Selector `json:"selector"`
	BindKey  value.Value       `json:"bindKey"`
}

func (i GetEvent) Act(wd selenium.WebDriver) error {
	elem, err := findElement(wd, i.Selector.Type.TypeName(), i.Selector.Value.Value())
	if err != nil {
		return fmt.Errorf("get event: find element error: %v", err)
	}
	s := store.NewStore()
	t, err := elem.Text()
	if err != nil {
		return fmt.Errorf("get event: retrieve text error: %v", err)
	}

	s.Set(i.BindKey.String(), t)

	return nil
}
