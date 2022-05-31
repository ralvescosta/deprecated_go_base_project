package main

import "github.com/spf13/cobra"

func NewSeedersCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "seeders",
		Short: "GoLang Base Application Seeders Command",
		Run: func(cmd *cobra.Command, args []string) {
			exec()
		},
	}
}
