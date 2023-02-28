package action

import "github.com/haytty/karas/internal/model/action/action_event"

type Action struct {
	Name        string                   `json:"name"`
	ActionEvent action_event.ActionEvent `json:"action_event"`
}
