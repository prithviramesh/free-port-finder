package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var RootCmd = &cobra.Command {
	Use: "portfinder",
	Short: "finds free TCP ports available on your machine",
	Long: "portfinder uses Go's built in net library to find free available TCP ports on localhost",
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("what do I put here?")
    },
}