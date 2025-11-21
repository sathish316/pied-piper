package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sathish316/pied-piper/config"
	"path/filepath"
	"os"
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
			Path: filepath.Join(os.Getenv("HOME"), config.DEFAULT_CONFIG_DIR),
			File: config.DEFAULT_CONFIG_FILE,
		}
		if err := config.Init(); err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing config: %v\n", err)
			os.Exit(1)
		}
	},
}

var configLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List config",
	Long:  `List config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing config...")
		config := config.Config{
			Path: filepath.Join(os.Getenv("HOME"), config.DEFAULT_CONFIG_DIR),
			File: config.DEFAULT_CONFIG_FILE,
		}
		config.Load()
		config.PrettyPrint()
	},
}

func init() {
	// Add subcommands to config
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configLsCmd)
}

