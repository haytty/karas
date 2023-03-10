package pseudo

import (
	"github.com/haytty/karas/internal/model"
	"github.com/haytty/karas/internal/model/action"
	"github.com/haytty/karas/internal/model/action/event"
	"github.com/haytty/karas/internal/model/action/selector"
	"github.com/haytty/karas/internal/model/action/value"
)

const (
	URL_COMMAND   = "URL"
	INPUT_COMMAND = "INPUT"
	CLICK_COMMAND = "CLICK"
	GET_COMMAND   = "GET"
)

type PseudoStruct interface {
	Materialize(k *model.Karas)
}

func collectMetaInfo(inlineMeta []string) []Meta {
	var m []Meta

	for i := 0; i < len(inlineMeta); i++ {
		meta, offset := NewMetaFromInline(inlineMeta[i:])
		m = append(m, meta)
		i += offset
	}

	return m
}

type PseudoURL struct {
	Value string
}

func (p PseudoURL) Materialize(k *model.Karas) {
	k.Suite.URL = p.Value
}

func NewPseudoURL(value string) *PseudoURL {
	return &PseudoURL{Value: value}
}

type PseudoInput struct {
	Value string
	Meta  []Meta
}

func NewPseudoInput(value string, inlineMeta []string) *PseudoInput {
	m := collectMetaInfo(inlineMeta)

	return &PseudoInput{
		Value: value,
		Meta:  m,
	}
}

func (p PseudoInput) Materialize(k *model.Karas) {
	event := new(event.InputEvent)

	event.Value = *value.NewValue(p.Value)

	var eventName string
	for _, meta := range p.Meta {
		switch m := meta.(type) {
		case *Selector:
			event.Selector = *selector.NewSelector(m.Type, m.Value)
		case *EventName:
			eventName = m.Value
		}
	}

	act := action.Action{
		Name:   eventName,
		Method: "input",
		Event:  event,
	}

	k.Suite.Actions = append(k.Suite.Actions, act)
}

type PseudoClick struct {
	Meta []Meta
}

func (p PseudoClick) Materialize(k *model.Karas) {
	event := new(event.ClickEvent)

	var eventName string
	for _, meta := range p.Meta {
		switch m := meta.(type) {
		case *Selector:
			event.Selector = *selector.NewSelector(m.Type, m.Value)
		case *EventName:
			eventName = m.Value
		}
	}

	act := action.Action{
		Name:   eventName,
		Method: "click",
		Event:  event,
	}

	k.Suite.Actions = append(k.Suite.Actions, act)
}

func NewPseudoClick(inlineMeta []string) *PseudoClick {
	m := collectMetaInfo(inlineMeta)

	return &PseudoClick{Meta: m}
}

type PseudoGet struct {
	Meta []Meta
}

func (p PseudoGet) Materialize(k *model.Karas) {
	event := new(event.GetEvent)

	var eventName string
	for _, meta := range p.Meta {
		switch m := meta.(type) {
		case *Selector:
			event.Selector = *selector.NewSelector(m.Type, m.Value)
		case *EventName:
			eventName = m.Value
		case *BindKey:
			event.BindKey = *value.NewValue(m.Value)
		}
	}

	act := action.Action{
		Name:   eventName,
		Method: "get",
		Event:  event,
	}

	k.Suite.Actions = append(k.Suite.Actions, act)
}

func NewPseudoGet(inlineMeta []string) *PseudoGet {
	m := collectMetaInfo(inlineMeta)

	return &PseudoGet{Meta: m}
}

func NewPseudoStruct(command string, args ...string) PseudoStruct {
	switch command {
	case URL_COMMAND:
		return NewPseudoURL(args[0])
	case INPUT_COMMAND:
		return NewPseudoInput(args[0], args[1:])
	case CLICK_COMMAND:
		return NewPseudoClick(args[0:])
	case GET_COMMAND:
		return NewPseudoGet(args[0:])
	}
	return nil
}
