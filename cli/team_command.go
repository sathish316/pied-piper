package cli

import (
	"fmt"

	config "github.com/sathish316/pied-piper/config"
	"github.com/spf13/cobra"
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
		teamName, _ := cmd.Flags().GetString("name")
		configPath := config.TeamConfigPath{
			Path: config.GetTeamConfigDir(teamName),
			File: config.DEFAULT_CONFIG_FILE,
		}
		configHandler := config.TeamConfigYamlHandler{
			ConfigPath: configPath,
		}
		err := configHandler.Init()
		if err != nil {
			fmt.Println("Error initializing team config: ", err)
			return
		}
		fmt.Printf("Team config initialized at %s\n", configPath.GetConfigFilePath())
	},
}

var teamShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show team config",
	Long:  `Show team config`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("name")
		configPath := config.TeamConfigPath{
			Path: config.GetTeamConfigDir(teamName),
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
	// Create command flags
	teamCreateCmd.Flags().StringP("name", "n", "pied-piper", "Team name")
	// Show command flags
	teamShowCmd.Flags().StringP("name", "n", "pied-piper", "Team name")
	// Add subcommands to config
	teamCmd.AddCommand(teamCreateCmd)
	teamCmd.AddCommand(teamShowCmd)
}
