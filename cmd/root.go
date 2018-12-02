package cmd

import (
	"github.com/spf13/cobra"
	"github.com/prithviramesh/free-port-finder/app"
	"fmt"
)

var RootCmd = &cobra.Command {
	Use: "free-port-finder",
	Short: "finds free TCP ports available on your machine",
	Long: "free-port-finder uses Go's built in net library to find free available TCP ports on localhost",
	Run: func(cmd *cobra.Command, args []string) {
        if port, err := app.FindPort(); err != nil {
			//TODO: implement better error handling
			panic(err)
		} else{
			fmt.Println(port)
		}
    },
}

func init() {
	RootCmd.AddCommand(CheckPortCmd)
}