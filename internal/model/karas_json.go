package model

import (
	"fmt"

	"github.com/haytty/karas/internal/model/action"
)

type KarasJSON struct {
	Karas struct {
		Url     string          `json:"url"`
		Output  string          `json:"output"`
		Actions []action.Action `json:"actions"`
	} `json:"karas"`
}

func NewKarasJSON() *KarasJSON {
	fmt.Println("KarasJSON")
	return &KarasJSON{}
}

func (s *KarasJSON) Do() error {
	return nil
}
