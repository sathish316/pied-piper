package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sathish316/pied-piper/config"
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
		config := config.Config{
			Path: config.DEFAULT_CONFIG_DIR,
			File: config.DEFAULT_CONFIG_FILE,
		}
		config.Init()
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

