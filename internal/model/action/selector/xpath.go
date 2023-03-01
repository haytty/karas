package selector

type XPath string

func (t *XPath) TypeName() string {
	return "XPath"
}

type XPathValue string

func (t *XPathValue) Value() string {
	return string(*t)
}
