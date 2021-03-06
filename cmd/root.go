package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var RootCmd = &cobra.Command {
	Use: "free-port-finder",
	Short: "free-port-finder uses Go's built in net library to find free available TCP ports on localhost\n",
	Example: "free-port-finder check <port-number> returns if the passed port is free\nfree-port-finder find returns a free port number if available\n",
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf(cmd.Example)
    },
}

func init() {
	RootCmd.AddCommand(CheckPortCmd)
	RootCmd.AddCommand(FindPortCmd)
}