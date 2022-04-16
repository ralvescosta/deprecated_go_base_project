/*
Copyright © 2022 Rafael Costa <rafael.rac.mg@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "base",
	Short: "GoLang base project",
	// Long:    `This project will be used to build other projects`,
	Version: "0.0.1",
}

func Execute(commands ...*cobra.Command) error {
	rootCmd.AddCommand(commands...)
	return rootCmd.Execute()
}
