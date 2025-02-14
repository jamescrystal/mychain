package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the main command for the blockchain CLI
var RootCmd = &cobra.Command{
	Use:   "mychaind",
	Short: "CLI for interacting with mychain blockchain",
}

func Execute() error {
	return RootCmd.Execute()
}
