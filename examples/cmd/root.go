package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sl",
		Short: "An example CLI application for working with IBM Classic Infrastructure",
		Long: `A collection of example commands that can be easily used to better understand how
the SoftLayer API works in golang. Authentication is set with the SL_USERNAME and SL_API_KEY env variables, or the
~/.softlayer config file.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

}

func initConfig() {

}
