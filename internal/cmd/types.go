package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sreallymatt/changeloggy/internal/config"
)

func NewTypesCommand(configPath *string) *cobra.Command {
	return &cobra.Command{
		Use:   "types",
		Short: "Lists all available entry types.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadAndValidate(*configPath)
			if err != nil {
				return err
			}

			for _, k := range cfg.Kinds {
				heading := "(no heading)"
				if k.Heading != nil {
					heading = *k.Heading
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%s -> %s\n", k.Name, heading)

				if len(k.Types) == 0 {
					fmt.Fprintf(cmd.OutOrStdout(), "  (no types defined)\n\n")
					continue
				}

				for _, t := range k.Types {
					line := strings.Builder{}
					fmt.Fprintf(&line, "  %-24s", t.Name)
					if t.Example != nil {
						fmt.Fprintf(&line, "  e.g. %s", *t.Example)
					}
					fmt.Fprintln(cmd.OutOrStdout(), line.String())
				}
				fmt.Fprintln(cmd.OutOrStdout())
			}
			return nil
		},
	}
}
