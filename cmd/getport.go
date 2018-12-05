package cmd

import (
	"os"
	"fmt"
	"github.com/prithviramesh/free-port-finder/app"
	"github.com/spf13/cobra"
)

var FindPortCmd = &cobra.Command {
	Use: "find",
	Short: "finds free TCP ports available on your machine",
	Long: "free-port-finder uses Go's built in net library to find free available TCP ports on localhost",
	Run: func(cmd *cobra.Command, args []string) {
        if port, err := app.FindPort(); err != nil {
			//TODO: implement better error handling
			fmt.Println("No free ports seem to exist")
			os.Exit(-1)
		} else{
			fmt.Println(port)
		}
    },
}