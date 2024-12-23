package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	// debugCmd represents the base command when called without any subcommands
	debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "Run the Plugin in debug mode",
		Run: func(cmd *cobra.Command, args []string) {
			err := Vessel.Mobilize()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	runCmd.AddCommand(debugCmd) // This enables the debug command for use while working on your Plugin
}
