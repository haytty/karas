package selector

import (
	"strings"

	"github.com/tebeka/selenium"
)

// FIX ME
// If xpath is specified, it may not work well. Therefore, it is recommended to work with css selector.
type XPath string

func (t *XPath) TypeName() string {
	return selenium.ByXPATH
}

type XPathValue string

func (t *XPathValue) Value() string {
	v := string(*t)
	trimedSuffix := strings.TrimSuffix(v, "\n")
	return strings.TrimSpace(trimedSuffix)
}
