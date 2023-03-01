package karas

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/haytty/karas/internal/model"
)

func Apply(jsonFile string) error {
	b, err := os.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	m := model.NewKarasJSON()
	if err := json.Unmarshal(b, m); err != nil {
		return err
	}

	fmt.Println(*m)

	fmt.Println("checkAction")
	for _, action := range m.Karas.Actions {
		action.Event.Act()
	}

	m.Do()

	return nil
}
