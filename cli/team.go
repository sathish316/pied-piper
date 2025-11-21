package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Manage team",
	Long:  `Manage team`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing Team command")
		// Add your team logic here
	},
}

