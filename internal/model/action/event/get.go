package event

import (
	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
)

type GetEvent struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}

func (i GetEvent) Act() error {
	return nil
}
