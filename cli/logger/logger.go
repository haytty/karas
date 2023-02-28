package logger

import (
	"github.com/haytty/golang_cli_template/cli/cli"
	"github.com/sirupsen/logrus"
)

func SetupLogger(c cli.Cli) error {
	logrus.SetOutput(c.Out())
	logrus.SetFormatter(&logrus.TextFormatter{})

	err := setLogLevel(c)

	return err
}

func setLogLevel(c cli.Cli) error {
	lvl, err := logrus.ParseLevel(c.LogLevel())
	if err != nil {
		return err
	}

	logrus.SetLevel(lvl)

	return nil
}
