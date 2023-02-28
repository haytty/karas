package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/haytty/golang_cli_template/cli/version"
	"github.com/haytty/golang_cli_template/internal/util"
)

type cliOption func(*GolangCliTemplateCli) error

func SetApplicationMode(cli *GolangCliTemplateCli) error {
	cli.appMode = os.Getenv("APP_ENV")

	return nil
}

func SetLogLevel(cli *GolangCliTemplateCli) error {
	validLogLevels := []string{
		"panic",
		"fatal",
		"error",
		"warn",
		"warning",
		"info",
		"debug",
		"trace",
	}
	loglevel := os.Getenv(fmt.Sprintf("%s_LOG_LEVEL", strings.ToUpper(version.Name)))
	cli.loglevel = loglevel

	if !util.SliceContains[string](validLogLevels, loglevel) {
		cli.loglevel = "info"
	}

	return nil
}

//func SetConfigDir(cli *GolangCliTemplateCli) error {
//	cli.configDir = os.Getenv(fmt.Sprintf("%s_CONFIG_DIR", strings.ToUpper(version.Name)))
//	return nil
//}
