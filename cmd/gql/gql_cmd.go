package gql

import "github.com/spf13/cobra"

func NewGraphQLCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gql",
		Short: "GoLang Base Application GraphQL Server Command",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}
