package pseudo

type Meta interface {
	TypeName() string
}

const (
	BINDKEY_TYPE   = "BIND_KEY"
	EVENTNAME_TYPE = "EVENT_NAME"
	SELECTOR_TYPE  = "SELECTOR"
)

type BindKey struct {
	Value string
}

func NewBindKey(value string) *BindKey {
	return &BindKey{Value: value}
}

func (s *BindKey) TypeName() string {
	return BINDKEY_TYPE
}

type EventName struct {
	Value string
}

func NewEventName(value string) *EventName {
	return &EventName{Value: value}
}

func (s *EventName) TypeName() string {
	return EVENTNAME_TYPE
}

type Selector struct {
	Type  string
	Value string
}

func (s *Selector) TypeName() string {
	return SELECTOR_TYPE
}

func NewSelector(Type string, value string) *Selector {
	return &Selector{Type: Type, Value: value}
}

func NewMetaFromInline(meta []string) (Meta, int) {
	switch meta[0] {
	case BINDKEY_TYPE:
		return NewBindKey(meta[1]), 1
	case SELECTOR_TYPE:
		return NewSelector(meta[1], meta[2]), 2
	case EVENTNAME_TYPE:
		return NewEventName(meta[1]), 1
	}

	return nil, 1
}
