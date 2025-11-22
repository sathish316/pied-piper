package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pied-piper",
	Short: "Pied-Piper - A powerful CLI tool for managing a team of AI SubAgents",
	Long:  `Pied-Piper - A powerful CLI tool for managing a team of AI SubAgents`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Add all commands to root
	rootCmd.AddCommand(teamCmd)
	rootCmd.AddCommand(subagentCmd)
}
