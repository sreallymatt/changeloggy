package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sreallymatt/changeloggy/internal/cmd"
	"github.com/sreallymatt/changeloggy/pkg/config"
)

//go:embed VERSION
var cliVersion string

func main() {
	var configPath string

	root := &cobra.Command{
		Use:           "changeloggy",
		Short:         "Changelog management with optional entry format validation",
		Version:       strings.TrimSpace(cliVersion), // TODO: remove this?
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.PersistentFlags().StringVarP(&configPath, "config", "c", config.FileName, "Path to the changeloggy configuration file.")

	// TODO: should this be located in `cmd/`?
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration management commands",
	}

	configCmd.AddCommand(cmd.NewConfigInitCommand(&configPath))
	configCmd.AddCommand(cmd.NewConfigValidateCommand(&configPath))

	root.AddCommand(configCmd)
	root.AddCommand(cmd.NewAddCommand(&configPath))
	root.AddCommand(cmd.NewCheckCommand(&configPath))
	root.AddCommand(cmd.NewGenerateCommand(&configPath))
	root.AddCommand(cmd.NewTypesCommand(&configPath))

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
