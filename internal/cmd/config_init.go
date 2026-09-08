package cmd

import (
	_ "embed"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sreallymatt/changeloggy/pkg/config"
)

//go:embed files/default_config.hcl
var defaultConfigContent []byte

func NewConfigInitCommand(configPath *string) *cobra.Command {
	var force bool

	c := &cobra.Command{
		Use:   "init",
		Short: "Generates a new changeloggy configuration file.",
		Long:  `Generates a new changeloggy configuration file with default and example values.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			path := *configPath
			if path == "" {
				path = config.FileName
			}

			_, err := os.Stat(path)
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("checking for existence of configuration file (%s): %w", path, err)
			}

			if err == nil && !force {
				return fmt.Errorf("configuration file (%s) already exists, to overwrite it, use `--force`", path)
			}

			if err := os.WriteFile(path, defaultConfigContent, 0o600); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "created default config at `%s`\n", path)
			return nil
		},
	}

	c.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing configuration file.")

	return c
}
