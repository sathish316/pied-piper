package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subagentCmd = &cobra.Command{
	Use:     "subagent",
	Aliases: []string{"subagents"},
	Short:   "Manage Subagents",
	Long:    `Manage Subagents`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing Subagent command")
		// Add your subagent logic here
	},
}

