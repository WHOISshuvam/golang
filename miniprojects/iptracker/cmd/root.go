package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Ip Tracker CLI app",
		Long:  `Ip Tracker CLI app`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
