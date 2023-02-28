package action_event

import (
	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
)

type Click struct {
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}
