package cli

import (
	"fmt"

	config "github.com/sathish316/pied-piper/config"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"github.com/sathish316/pied-piper/generator"
)

var subagentCmd = &cobra.Command{
	Use:     "subagent",
	Aliases: []string{"subagents"},
	Short:   "Manage Subagents",
	Long:    `Manage Subagents`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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

var subagentGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate subagent for a team and Coding Agent (Claude Code, Cursor, etc.)",
	Long:  `Generate subagent for a team and Coding Agent (Claude Code, Cursor, etc.)`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("team")
		subagentName, _ := cmd.Flags().GetString("name")
		target, _ := cmd.Flags().GetString("target")
		projectDir, _ := cmd.Flags().GetString("projectDir")
		var codingAgentTarget config.CodingAgentConfig
		// Initialize Coding Agent (Claude Code) specific config
		if target == "claude-code" {
			claudeCodingAgent := &config.ClaudeCodingAgent{}
			if projectDir == "" {
				targetDir := claudeCodingAgent.GetUserSubagentConfigDir()
				fmt.Printf("Generating subagent %s for team %s and target %s in User directory %s...\n", subagentName, teamName, target, targetDir)
				codingAgentTarget = config.CodingAgentConfig{
					Target: config.ClaudeCode,
					TargetDir: targetDir,
					TargetDirType: config.TargetDirTypeUser,
				}
			} else {
				targetDir := claudeCodingAgent.GetProjectSubagentConfigDir(projectDir)
				fmt.Printf("Generating subagent %s for team %s and target %s in Project directory %s...\n", subagentName, teamName, target, targetDir)
				codingAgentTarget = config.CodingAgentConfig{
					Target: config.ClaudeCode,
					TargetDir: targetDir,
					TargetDirType: config.TargetDirTypeProject,
				}
			}
		} else {
			fmt.Printf("Target %s is not supported. Only claude-code is supported currently.\n", target)
			return
		}
		// Get team config
		teamConfig, err := getTeamConfig()
		if err != nil {
			fmt.Println("Error getting team config: ", err)
			return
		}
		// Get subagent config
		subagentConfig, err := teamConfig.FindSubagentByRole(subagentName)
		if err != nil {
			fmt.Println("Error getting subagent config: ", err)
			return
		}
		// Generate subagent yaml for coding agent
		subagentGenerator := &generator.SDLCSubAgentGenerator{
			TeamConfig: teamConfig,
		}
		subagentYamlFilePath, err := subagentGenerator.GenerateSubagentYamlForCodingAgent(subagentConfig, codingAgentTarget)
		if err != nil {
			fmt.Println("Error generating Subagent for %s:", err, codingAgentTarget.ToString())
			return
		}
		fmt.Printf("Subagent yaml file generated at %s\n", subagentYamlFilePath)
	},
}

// FIXME: Make this work for multiple teams
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
	// List config - flags, default,required
	subagentListCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	// Show config - flags, default,required
	subagentShowCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentShowCmd.Flags().StringP("name", "s", "", "Subagent name")
	subagentShowCmd.MarkFlagRequired("name")
	// Generate config - flags, default, required
	subagentGenerateCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentGenerateCmd.Flags().StringP("name", "s", "", "Subagent name")
	subagentGenerateCmd.Flags().StringP("target", "f", "", "Target coding agent (claude-code, cursor, etc.). Only claude-code is supported currently.")
	subagentGenerateCmd.Flags().StringP("projectDir", "p", "", "Subagents are generated for a specific project directory. If not provided, subagents are generated in User directory for target.")
	subagentGenerateCmd.MarkFlagRequired("name")
	subagentGenerateCmd.MarkFlagRequired("target")
	// Add sub-commands
	subagentCmd.AddCommand(subagentListCmd)
	subagentCmd.AddCommand(subagentShowCmd)
	subagentCmd.AddCommand(subagentGenerateCmd)
}
