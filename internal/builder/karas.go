package builder

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"github.com/haytty/karas/internal/parser/pseudo"
	"github.com/haytty/karas/internal/util"

	"github.com/haytty/karas/internal/model"

	"github.com/haytty/karas/internal/webdriver"
)

func BuildFromKarasJSON(jsonFilePath string, selenium, chrome, chromeDriver string, port int) (*model.Karas, error) {
	w := webdriver.NewSelenium(selenium, chrome, chromeDriver, port)

	karas := model.NewKaras(w)

	b, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, karas); err != nil {
		return nil, err
	}

	return karas, nil
}

func BuildFromConfig(configFilePath string, selenium, chrome, chromeDriver string, port int) (*model.Karas, error) {
	w := webdriver.NewSelenium(selenium, chrome, chromeDriver, port)

	karas := model.NewKaras(w)

	lines, err := util.ReadLines(configFilePath)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		var fields []string

		reader := strings.NewReader(line)
		scanner := bufio.NewScanner(reader)

		scanner.Split(util.ScanWordsForQuotes)

		for scanner.Scan() {
			fields = append(fields, scanner.Text())
		}

		pstruct := pseudo.NewPseudoStruct(fields[0], fields[1:]...)

		pstruct.Materialize(karas)
	}

	return karas, nil
}
