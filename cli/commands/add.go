package commands

import (
	"github.com/haytty/karas/cli/cli"
	karas "github.com/haytty/karas/internal/handler/karas/add"
	"github.com/spf13/cobra"
)

func AddCommand(cli cli.Cli) *cobra.Command {
	acceptArglength := 2

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "This is short message.",
		Long:  "This is long message.",
		Args:  cobra.MatchAll(cobra.ExactArgs(acceptArglength)),
		RunE: func(cmd *cobra.Command, args []string) error {
			param1 := args[0]
			param2 := args[1]
			if err := karas.Apply(param1, param2); err != nil {
				return err
			}

			return nil
		},
	}

	return addCmd
}
