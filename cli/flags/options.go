package flags

import "github.com/haytty/karas/internal/util"

type GlobalOption struct {
	Config string
	Json   string
}

func (g *GlobalOption) Valid() bool {
	return g.validConfig()
}

func (g *GlobalOption) validConfig() bool {
	return util.IsFileExist(g.Config) || (g.Json != "" && util.IsFileExist(g.Json))
}

var globalOption *GlobalOption

func NewGlobalOption() *GlobalOption {
	if globalOption == nil {
		globalOption = &GlobalOption{}
	}

	return globalOption
}
