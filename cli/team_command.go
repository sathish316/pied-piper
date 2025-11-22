package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	team "github.com/sathish316/pied-piper/team"
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
		teamConfig := team.TeamConfig{
			Path: filepath.Join(os.Getenv("HOME"), team.DEFAULT_CONFIG_DIR),
			File: team.DEFAULT_CONFIG_FILE,
		}
		teamObj := team.Team{
			TeamConfig: &teamConfig,
		}
		configHandler := &team.TeamConfigYamlHandler{
			Team: &teamObj,
		}
		configHandler.Init()
	},
}

var teamShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show team config",
	Long:  `Show team config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Showing team config...")
		teamConfig := team.TeamConfig{
			Path: filepath.Join(os.Getenv("HOME"), team.DEFAULT_CONFIG_DIR),
			File: team.DEFAULT_CONFIG_FILE,
		}
		teamObj := team.Team{
			TeamConfig: &teamConfig,
		}
		configHandler := team.TeamConfigYamlHandler{
			Team: &teamObj,
		}
		configHandler.Load()
		configHandler.PrettyPrint()
	},
}

func init() {
	// Add subcommands to config
	teamCmd.AddCommand(teamCreateCmd)
	teamCmd.AddCommand(teamShowCmd)
}

