package selector

import (
	"strings"

	"github.com/tebeka/selenium"
)

type CSS string

func (t *CSS) TypeName() string {
	return selenium.ByCSSSelector
}

type CSSValue string

func (t *CSSValue) Value() string {
	v := string(*t)
	trimedSuffix := strings.TrimSuffix(v, "\n")
	return strings.TrimSpace(trimedSuffix)
}
