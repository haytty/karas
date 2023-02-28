package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/haytty/karas/cli/cli"
	"github.com/haytty/karas/cli/commands"
	"github.com/haytty/karas/cli/flags"
	"github.com/haytty/karas/cli/logger"
	"github.com/haytty/karas/cli/version"
	"github.com/haytty/karas/internal/handler/karas"
	"github.com/spf13/cobra"
)

func NewKarasCommand(cli cli.Cli) *cobra.Command {
	rootCmd := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "karas",
		Short: "This is short message.",
		Long: fmt.Sprintln(
			"This is long message.\n" +
				"This is long message..\n" +
				"This is long message."),
		Version:       version.CurrentVersion(),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return karas.Apply()
		},
		PersistentPreRunE: initialize(cli),
	}
	rootCmd.AddCommand(
		commands.AddCommand(cli),
	)

	opts := flags.NewGlobalOption()
	flagName := "base-dir"
	defaultDir := filepath.Join(os.Getenv("HOME"), ".config", "karas")
	rootCmd.PersistentFlags().StringVarP(
		&opts.BaseDir,
		flagName,
		"d",
		defaultDir,
		"base directory",
	)

	if err := rootCmd.RegisterFlagCompletionFunc(
		flagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{defaultDir}, cobra.ShellCompDirectiveFilterFileExt
		}); err != nil {
		os.Exit(1)
	}

	return rootCmd
}

func initialize(c cli.Cli) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// opts := flags.NewGlobalOption()

		if err := logger.SetupLogger(c); err != nil {
			return fmt.Errorf("setup logger: %w", err)
		}

		return nil
	}
}
