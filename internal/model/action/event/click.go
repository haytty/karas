package event

import (
	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
)

type ClickEvent struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}

func (i ClickEvent) Act() error {
	return nil
}
