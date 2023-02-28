package model

import "github.com/haytty/karas/internal/model/action"

type KarasJSON struct {
	Url     string          `json:"url"`
	Output  string          `json:"output"`
	Actions []action.Action `json:"action"`
}

func NewKarasJSON() *KarasJSON {
	return &KarasJSON{}
}
