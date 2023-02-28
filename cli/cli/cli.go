package cli

import (
	"fmt"
	"io"
	"os"
)

type Cli interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer

	LogLevel() string
}

type DefaultCli struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer

	loglevel string
}

func (cli *DefaultCli) In() io.Reader {
	return cli.stdin
}

func (cli *DefaultCli) Out() io.Writer {
	return cli.stdout
}

func (cli *DefaultCli) Err() io.Writer {
	return cli.stderr
}

func (cli *DefaultCli) LogLevel() string {
	return cli.loglevel
}

type KarasCli struct {
	DefaultCli

	appMode string

	// configDir string
}

func (cli *KarasCli) Apply(opts ...cliOption) error {
	for _, opt := range opts {
		err := opt(cli)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewKarasCli(opts ...cliOption) *KarasCli {
	cli := &KarasCli{
		DefaultCli: DefaultCli{
			stdin:  os.Stdin,
			stdout: os.Stdout,
			stderr: os.Stderr,
		},
	}

	defaultOptions := []cliOption{
		SetApplicationMode,
		SetLogLevel,
	}

	opts = append(opts, defaultOptions...)

	err := cli.Apply(opts...)
	if err != nil {
		fmt.Fprintln(cli.stderr, err)
		os.Exit(1)
	}

	return cli
}
