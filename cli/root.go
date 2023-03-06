package cli

import (
	"fmt"
	"path/filepath"

	"github.com/haytty/karas/cli/cli"
	"github.com/haytty/karas/cli/flags"
	"github.com/haytty/karas/cli/logger"
	"github.com/haytty/karas/cli/version"
	"github.com/haytty/karas/internal/handler/karas"
	"github.com/spf13/cobra"
)

func NewKarasCommand(cli cli.Cli) *cobra.Command {
	opts := flags.NewGlobalOption()

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
			return karas.Apply(opts)
		},
		PersistentPreRunE: initialize(cli),
	}

	configFlagName := "config"
	defaultKarasfile := filepath.Join("./", "Karasfile")
	rootCmd.PersistentFlags().StringVarP(&opts.Config, configFlagName, "c", defaultKarasfile, "config file path")

	chromeDriverFlagName := "chrome-driver"
	rootCmd.PersistentFlags().StringVarP(&opts.ChromeDriver, chromeDriverFlagName, "", "chromedriver", "chrome driver path")

	chromeFlagName := "chrome"
	rootCmd.PersistentFlags().StringVarP(&opts.Chrome, chromeFlagName, "", "chrome", "chrome binary path")

	seleniumFlagName := "selenium"
	rootCmd.PersistentFlags().StringVarP(&opts.SeleniumPath, seleniumFlagName, "", "selenium-server.jar", "selenium server path")

	jsonFlagName := "json"
	rootCmd.PersistentFlags().StringVarP(&opts.JSON, jsonFlagName, "j", "", "json file path")

	portFlagName := "port"
	rootCmd.PersistentFlags().IntVarP(&opts.Port, portFlagName, "p", 8080, "selenium localize port.")

	return rootCmd
}

func initialize(c cli.Cli) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		opts := flags.NewGlobalOption()
		if !opts.Valid() {
			return fmt.Errorf("%s is not found.", opts.Config)
		}

		if err := logger.SetupLogger(c); err != nil {
			return fmt.Errorf("setup logger: %w", err)
		}

		return nil
	}
}
