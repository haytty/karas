package action

import (
	"encoding/json"

	"github.com/haytty/karas/internal/model/action/event"
)

type Action struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Event  event.Event
}

func (a *Action) UnmarshalJSON(bytes []byte) error {
	preUnMarshalAction := struct {
		Name   string          `json:"name"`
		Method string          `json:"method"`
		Event  json.RawMessage `json:"event"`
	}{}

	if err := json.Unmarshal(bytes, &preUnMarshalAction); err != nil {
		return err
	}

	switch preUnMarshalAction.Method {
	case "input":
		event := new(event.InputEvent)
		if err := json.Unmarshal(preUnMarshalAction.Event, event); err != nil {
			return err
		}
		a.Event = event
	case "click":
		event := new(event.ClickEvent)
		if err := json.Unmarshal(preUnMarshalAction.Event, event); err != nil {
			return err
		}
		a.Event = event
	case "get":
		event := new(event.GetEvent)
		if err := json.Unmarshal(preUnMarshalAction.Event, event); err != nil {
			return err
		}
		a.Event = event
	case "submit":
		event := new(event.SubmitEvent)
		if err := json.Unmarshal(preUnMarshalAction.Event, event); err != nil {
			return err
		}
		a.Event = event
	}

	a.Name = preUnMarshalAction.Name
	a.Method = preUnMarshalAction.Method

	return nil
}
