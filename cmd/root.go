package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/privateerproj/privateer-sdk/command"
	"github.com/privateerproj/privateer-sdk/shared"
)

var (
	// Build information is added by the Makefile at compile time
	buildVersion       string
	buildGitCommitHash string
	buildTime          string

	PluginName = "osps-baseline"

	// runCmd represents the base command when called without any subcommands
	runCmd = &cobra.Command{
		Use:   PluginName,
		Short: fmt.Sprintf("Test suite for %s.", PluginName),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// optional
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Serve plugin
			plugin := &Plugin{}
			serveOpts := &shared.ServeOpts{
				Plugin: plugin,
			}

			shared.Serve(PluginName, serveOpts)
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the runCmd.
func Execute(version, commitHash, builtAt string) {
	buildVersion = version
	buildGitCommitHash = commitHash
	buildTime = builtAt

	err := runCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	command.SetBase(runCmd) // This initializes the base CLI functionality
}
