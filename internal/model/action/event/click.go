package payload

import (
	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
)

type ClickPayload struct {
	Payload
	Selector selector.Selector `json:"selector"`
	Value    value.Value       `json:"value"`
}
