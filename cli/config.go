package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage config",
	Long:  `Manage config`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config",
	Long:  `Initialize config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing config...")
		// Add your initialization logic here
	},
}

var configLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List config",
	Long:  `List config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing config...")
		// Add your list logic here
	},
}

func init() {
	// Add subcommands to config
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configLsCmd)
}

