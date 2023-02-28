package command

import (
	"fmt"
	"os"
	"os/exec"
)

type Command struct {
	command *exec.Cmd
}

func NewCommand(name string, arg ...string) *Command {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return &Command{
		command: cmd,
	}
}

func (c *Command) Execute() error {
	if err := c.command.Run(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
