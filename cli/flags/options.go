package flags

type GlobalOption struct {
	BaseDir string
}

var globalOption *GlobalOption

func NewGlobalOption() *GlobalOption {
	if globalOption == nil {
		globalOption = &GlobalOption{}
	}

	return globalOption
}
