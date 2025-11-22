package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	config "github.com/sathish316/pied-piper/config"
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
		teamConfig, err := getTeamConfig()
		if err != nil {
			fmt.Println("Error getting team config: ", err)
			return
		}
		subagentConfigHandler := &config.SubagentConfigYamlHandler{
			Config: teamConfig,
		}
		subagents, err := subagentConfigHandler.List(teamName)
		if err != nil {
			fmt.Println("Error listing subagents: ", err)
			return
		}
		for _, subagent := range subagents {
			fmt.Printf("Subagent: %s\n", subagent.Role)
		}
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
		teamConfig, err := getTeamConfig()
		if err != nil {
			fmt.Println("Error getting team config: ", err)
			return
		}
		subagentConfigHandler := &config.SubagentConfigYamlHandler{
			Config: teamConfig,
		}
		subagent, err := subagentConfigHandler.Show(teamName, subagentName)
		if err != nil {
			fmt.Println("Error showing subagent: ", err)
			return
		}
		fmt.Printf("Subagent: %s\n", subagent.ToString())
	},
}

//FIXME: Make this work for multiple teams
func getTeamConfig() (*config.TeamConfig, error) {
	configPath := config.TeamConfigPath{
		Path: filepath.Join(os.Getenv("HOME"), config.DEFAULT_CONFIG_DIR),
		File: config.DEFAULT_CONFIG_FILE,
	}
	configHandler := config.TeamConfigYamlHandler{
		ConfigPath: configPath,
	}
	teamConfig, err := configHandler.Load()
	if err != nil {
		fmt.Println("Error loading team config: ", err)
		return nil, err
	}
	return teamConfig, nil
}

func init() {
	subagentListCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentShowCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentShowCmd.Flags().StringP("name", "s", "", "Subagent name")
	subagentShowCmd.MarkFlagRequired("name")
	subagentCmd.AddCommand(subagentListCmd)
	subagentCmd.AddCommand(subagentShowCmd)
}