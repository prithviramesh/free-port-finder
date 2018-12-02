package cmd

import (
	"os"
	"fmt"
	"strconv"
	"github.com/prithviramesh/free-port-finder/app"
	"github.com/spf13/cobra"
)

var CheckPortCmd = &cobra.Command {
	Use: "check port",
	Short: "checks if a given port is free",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
		port := args[0]

		s, err := strconv.Atoi(port)
		if err != nil || (s < 0 || s > 65535) {
			fmt.Println("Enter a valid port value between the range 0-65535")
			os.Exit(-1)
		}

		if check, _ := app.CheckPort(port); !check {
			fmt.Printf("Port %s is in use\n", port)
		} else {
			fmt.Printf("Port %s is free to use\n", port)
		}
	},
}