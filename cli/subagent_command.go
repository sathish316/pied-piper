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

var subagentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List subagents of a team",
	Long:  `List subagents of a team`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("team")
		fmt.Printf("Listing subagents of team %s...\n", teamName)
	},
}

var subagentShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show subagent of a team",
	Long:  `Show subagent of a team`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("team")
		subagentName, _ := cmd.Flags().GetString("name")
		fmt.Printf("Showing subagent %s of team %s...\n", subagentName, teamName)
	},
}

func init() {
	subagentListCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentShowCmd.Flags().StringP("name", "s", "", "Subagent name")
	subagentShowCmd.MarkFlagRequired("name")
	subagentCmd.AddCommand(subagentListCmd)
	subagentCmd.AddCommand(subagentShowCmd)
}