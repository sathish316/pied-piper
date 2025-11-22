package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	config "github.com/sathish316/pied-piper/config"
)

var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Manage team",
	Long:  `Manage team`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var teamCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create team and initialize config",
	Long:  `Create team and initialize config`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.TeamConfigPath{
			Path: filepath.Join(os.Getenv("HOME"), config.DEFAULT_CONFIG_DIR),
			File: config.DEFAULT_CONFIG_FILE,
		}
		configHandler := &config.TeamConfigYamlHandler{
			ConfigPath: configPath,
		}
		configHandler.Init()
	},
}

var teamShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show team config",
	Long:  `Show team config`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.TeamConfigPath{
			Path: filepath.Join(os.Getenv("HOME"), config.DEFAULT_CONFIG_DIR),
			File: config.DEFAULT_CONFIG_FILE,
		}
		fmt.Println("Showing team config from file ", configPath.GetConfigFilePath())
		configHandler := config.TeamConfigYamlHandler{
			ConfigPath: configPath,
		}
		_, err := configHandler.Load()
		if err != nil {
			fmt.Println("Error loading team config: ", err)
			return
		}
		configStr, err := configHandler.PrettyPrint()
		if err != nil {
			fmt.Println("Error pretty printing team config: ", err)
			return
		}
		fmt.Println(configStr)
	},
}

func init() {
	// Add subcommands to config
	teamCmd.AddCommand(teamCreateCmd)
	teamCmd.AddCommand(teamShowCmd)
}

