package flags

import "github.com/haytty/karas/internal/util"

type GlobalOption struct {
	Config string
	JSON   string

	Format string

	ChromeDriver string
	Chrome       string

	SeleniumPath string
	Port         int
}

func (g *GlobalOption) Valid() bool {
	return g.validConfig()
}

func (g *GlobalOption) validConfig() bool {
	return util.IsFileExist(g.Config) || (g.JSON != "" && util.IsFileExist(g.JSON))
}

var globalOption *GlobalOption

func NewGlobalOption() *GlobalOption {
	if globalOption == nil {
		globalOption = &GlobalOption{}
	}

	return globalOption
}
