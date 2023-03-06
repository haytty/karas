package selector

import (
	"encoding/json"
)

type Selector struct {
	Type  SelectorType  `json:"type"`
	Value SelectorValue `json:"value"`
}

func (s *Selector) UnmarshalJSON(bytes []byte) error {
	preUnMarshalSelector := struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}{}

	if err := json.Unmarshal(bytes, &preUnMarshalSelector); err != nil {
		return err
	}

	switch preUnMarshalSelector.Type {
	case "xpath":
		t, v :=
			XPath(preUnMarshalSelector.Type),
			XPathValue(preUnMarshalSelector.Value)

		s.Type = &t
		s.Value = &v
	case "css":
		t, v :=
			CSS(preUnMarshalSelector.Type),
			CSSValue(preUnMarshalSelector.Value)

		s.Type = &t
		s.Value = &v
	}
	return nil
}

type SelectorType interface {
	TypeName() string
}

type SelectorValue interface {
	Value() string
}
