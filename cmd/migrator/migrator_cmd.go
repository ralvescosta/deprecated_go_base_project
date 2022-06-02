package migrator

import "github.com/spf13/cobra"

func NewMigratorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "seeders",
		Short: "GoLang Base Application Migration Command",
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}
}
