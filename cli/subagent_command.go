package cli

import (
	"fmt"

	"os"
	"path/filepath"

	config "github.com/sathish316/pied-piper/config"
	"github.com/sathish316/pied-piper/generator"
	"github.com/spf13/cobra"
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
		teamConfig, err := getTeamConfig(teamName)
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
		teamConfig, err := getTeamConfig(teamName)
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

var subagentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new subagent in a team",
	Long:  `Create a new subagent in a team with placeholders for required fields`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("team")
		role, _ := cmd.Flags().GetString("role")
		nickname, _ := cmd.Flags().GetString("nickname")

		fmt.Printf("Creating subagent '%s' for team '%s'...\n", role, teamName)

		// Load team config
		teamConfig, err := getTeamConfig(teamName)
		if err != nil {
			fmt.Println("Error getting team config: ", err)
			return
		}

		// Check if subagent already exists
		_, err = teamConfig.FindSubagentByRole(role)
		if err == nil {
			fmt.Printf("Error: Subagent with role '%s' already exists in team '%s'\n", role, teamName)
			return
		}

		// Create new subagent with empty placeholders
		newSubagent := config.SubagentConfig{
			Role:        role,
			Description: "",
			Nickname:    nickname,
			TaskLabels: config.TaskLabelsConfig{
				Incoming:                []string{},
				Outgoing:                []string{},
				TaskWorkflowDescription: "",
			},
			WikiLabels: config.WikiLabelsConfig{
				Incoming:                []string{},
				Outgoing:                []string{},
				WikiWorkflowDescription: "",
			},
		}

		// Add to team config
		teamConfig.SubAgents = append(teamConfig.SubAgents, newSubagent)

		// Save updated config
		configPath := config.TeamConfigPath{
			Path: filepath.Join(os.Getenv("HOME"), config.DEFAULT_CONFIG_DIR, teamName),
			File: config.DEFAULT_CONFIG_FILE,
		}
		configHandler := config.TeamConfigYamlHandler{
			ConfigPath: configPath,
			Config:     teamConfig,
		}
		err = configHandler.Save()
		if err != nil {
			fmt.Println("Error saving team config: ", err)
			return
		}

		nicknameDisplay := "none"
		if nickname != "" {
			nicknameDisplay = nickname
		}
		fmt.Printf("\nSubagent '%s' (Nickname: %s) created successfully!\n", role, nicknameDisplay)
	},
}

var subagentGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate subagent for a team and Coding Agent (Claude Code, Rovodev, Cursor, etc.)",
	Long:  `Generate subagent for a team and Coding Agent (Claude Code, Rovodev, Cursor, etc.)`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("team")
		subagentName, _ := cmd.Flags().GetString("name")
		target, _ := cmd.Flags().GetString("target")
		projectDir, _ := cmd.Flags().GetString("projectDir")
		codingAgentTarget, err := getCodingAgentTarget(target, projectDir)
		if err != nil {
			fmt.Println("CodingAgentConfig Error: ", err)
			return
		}
		fmt.Printf("Generating subagent %s for team %s and Coding Agent target: %s...\n", subagentName, teamName, codingAgentTarget.ToString())
		// Get team config
		teamConfig, err := getTeamConfig(teamName)
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
		subagentSpecFilePath, err := subagentGenerator.GenerateSubagentSpecFileForCodingAgent(subagentConfig, codingAgentTarget)
		if err != nil {
			fmt.Printf("Error generating Subagent for %s: %s\n", codingAgentTarget.ToString(), err)
			return
		}
		fmt.Printf("Subagent spec file generated at %s\n", subagentSpecFilePath)
	},
}

var subagentGenerateAllCmd = &cobra.Command{
	Use:   "generate-all",
	Short: "Generate all subagents for a team and Coding Agent (Claude Code, Cursor, etc.)",
	Long:  `Generate all subagents for a team and Coding Agent (Claude Code, Cursor, etc.)`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("team")
		target, _ := cmd.Flags().GetString("target")
		projectDir, _ := cmd.Flags().GetString("projectDir")
		codingAgentTarget, err := getCodingAgentTarget(target, projectDir)
		if err != nil {
			fmt.Println("CodingAgentConfig Error: ", err)
			return
		}
		fmt.Printf("Generating all subagents for team %s and Coding Agent target: %s...\n", teamName, codingAgentTarget.ToString())
		// Get team config
		teamConfig, err := getTeamConfig(teamName)
		if err != nil {
			fmt.Println("Error getting team config: ", err)
			return
		}
		// Get all subagents config
		for _, subagentConfig := range teamConfig.SubAgents {
			// Generate subagent yaml for coding agent
			subagentGenerator := &generator.SDLCSubAgentGenerator{
				TeamConfig: teamConfig,
			}
			subagentSpecFilePath, err := subagentGenerator.GenerateSubagentSpecFileForCodingAgent(&subagentConfig, codingAgentTarget)
			if err != nil {
				fmt.Printf("Error generating Subagent for %s: %s\n", codingAgentTarget.ToString(), err)
				return
			}
			fmt.Printf("Subagent spec file generated at %s\n", subagentSpecFilePath)
		}
	},
}

var subagentGenerateMetapromptCmd = &cobra.Command{
	Use:   "metaprompt",
	Short: "Show metaprompt for a subagent",
	Long:  `Show metaprompt for a subagent`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Use the below prompt from cursor or Claude-code\n")
		fmt.Printf("Prompt: %s\n", config.GetSubagentMetapromptContent())
	},
}

// FIXME: Make this work for multiple teams
func getTeamConfig(teamName string) (*config.TeamConfig, error) {
	configPath := config.TeamConfigPath{
		Path: filepath.Join(os.Getenv("HOME"), config.DEFAULT_CONFIG_DIR, teamName),
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

func getCodingAgentTarget(codingAgent string, projectDir string) (*config.CodingAgentConfig, error) {
	// Initialize Coding Agent (Claude Code) specific config
	if codingAgent == string(config.ClaudeCode) {
		claudeCodingAgent := &config.ClaudeCodingAgent{}
		if projectDir == "" {
			targetDir := claudeCodingAgent.GetUserSubagentConfigDir()
			codingAgentTarget := &config.CodingAgentConfig{
				Target:        config.ClaudeCode,
				TargetDir:     targetDir,
				TargetDirType: config.TargetDirTypeUser,
			}
			return codingAgentTarget, nil
		} else {
			targetDir := claudeCodingAgent.GetProjectSubagentConfigDir(projectDir)
			codingAgentTarget := &config.CodingAgentConfig{
				Target:        config.ClaudeCode,
				TargetDir:     targetDir,
				TargetDirType: config.TargetDirTypeProject,
			}
			return codingAgentTarget, nil
		}
	} else if codingAgent == string(config.Rovodev) {
			rovodevCodingAgent := &config.RovodevCodingAgent{}
			if projectDir == "" {
				targetDir := rovodevCodingAgent.GetUserSubagentConfigDir()
				codingAgentTarget := &config.CodingAgentConfig{
					Target:        config.Rovodev,
					TargetDir:     targetDir,
					TargetDirType: config.TargetDirTypeUser,
				}
				return codingAgentTarget, nil
			} else {
				targetDir := rovodevCodingAgent.GetProjectSubagentConfigDir(projectDir)
				codingAgentTarget := &config.CodingAgentConfig{
					Target:        config.Rovodev,
					TargetDir:     targetDir,
					TargetDirType: config.TargetDirTypeProject,
				}
				return codingAgentTarget, nil
			}
	} else {
		return nil, fmt.Errorf("target %s is not supported", codingAgent)
	}
}

func init() {
	// List config - flags, default,required
	subagentListCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	// Show config - flags, default,required
	subagentShowCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentShowCmd.Flags().StringP("name", "s", "", "Subagent name")
	subagentShowCmd.MarkFlagRequired("name")
	// Create config - flags, default, required
	subagentCreateCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentCreateCmd.Flags().StringP("role", "r", "", "Subagent role (unique identifier)")
	subagentCreateCmd.Flags().StringP("nickname", "n", "", "Subagent nickname (optional)")
	subagentCreateCmd.MarkFlagRequired("team")
	subagentCreateCmd.MarkFlagRequired("role")
	// Generate config - flags, default, required
	subagentGenerateCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentGenerateCmd.Flags().StringP("name", "s", "", "Subagent name")
	subagentGenerateCmd.Flags().StringP("target", "f", "", "Target coding agent (claude-code, rovodev, cursor, etc.). Only claude-code is supported currently.")
	subagentGenerateCmd.Flags().StringP("project-dir", "p", "", "Subagents are generated for a specific project directory. If not provided, subagents are generated in User directory for target.")
	subagentGenerateCmd.Flags().BoolP("skip-llm", "l", false, "Skip LLM API keys for subagent generation. Get prompts that you can run from Cursor or Claude Code instead.")
	subagentGenerateCmd.MarkFlagRequired("name")
	subagentGenerateCmd.MarkFlagRequired("target")
	// Generate all config - flags, default, required
	subagentGenerateAllCmd.Flags().StringP("team", "t", "pied-piper", "Team name")
	subagentGenerateAllCmd.Flags().StringP("target", "f", "", "Target coding agent (claude-code, rovodev, cursor, etc.). Only claude-code is supported currently.")
	subagentGenerateAllCmd.Flags().StringP("project-dir", "p", "", "Subagents are generated for a specific project directory. If not provided, subagents are generated in User directory for target.")
	subagentGenerateAllCmd.Flags().BoolP("skip-llm", "l", false, "Skip LLM API keys for subagent generation. Get prompts that you can run from Cursor or Claude Code instead.")
	subagentGenerateAllCmd.MarkFlagRequired("target")
	// Add sub-commands
	subagentCmd.AddCommand(subagentListCmd)
	subagentCmd.AddCommand(subagentShowCmd)
	subagentCmd.AddCommand(subagentCreateCmd)
	subagentCmd.AddCommand(subagentGenerateCmd)
	subagentCmd.AddCommand(subagentGenerateAllCmd)
	subagentCmd.AddCommand(subagentGenerateMetapromptCmd)
}
