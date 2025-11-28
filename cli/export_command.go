package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sathish316/pied-piper/config"
	"github.com/sathish316/pied-piper/transporter"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export subagent config files",
	Long:  `Export subagent config files`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var exportAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Export all subagent config files",
	Long:  `Export all subagent config files`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("team")
		projectDir, _ := cmd.Flags().GetString("project-dir")
		target, _ := cmd.Flags().GetString("target")
		if target != "claude-code" {
			fmt.Printf("Error: unsupported target '%s'. Only 'claude-code' is currently supported.\n", target)
			return
		}
		teamConfig, err := getTeamConfig(teamName)
		if err != nil {
			fmt.Printf("Error getting team config: %s\n", err)
			return
		}
		exporter := &transporter.ClaudeCodeSubAgentExporter{
			TeamConfig: teamConfig,
			CodingAgent: &config.ClaudeCodingAgent{},
			FileUtils: &transporter.FileUtils{},
		}
		if projectDir == "" {
			fmt.Printf("Exporting all %s subagent config files to user directory\n", teamName)
			err := exporter.ExportAll()
			if err != nil {
				fmt.Printf("Error exporting all subagent config files to user directory: %s\n", err)
				return
			}
		} else {
			fmt.Printf("Exporting all %s subagent config files to project directory %s\n", teamName, projectDir)
			err := exporter.ExportAllToProject(projectDir)
			if err != nil {
				fmt.Printf("Error exporting all subagent config files to project directory: %s\n", err)
				return
			}
		}
	},
}

var exportSubagentCmd = &cobra.Command{
	Use:   "subagent",
	Short: "Export subagent config file",
	Long:  `Export subagent config file`,
	Run: func(cmd *cobra.Command, args []string) {
		teamName, _ := cmd.Flags().GetString("team")
		subagentName, _ := cmd.Flags().GetString("name")
		projectDir, _ := cmd.Flags().GetString("project-dir")
		target, _ := cmd.Flags().GetString("target")
		if target != "claude-code" {
			fmt.Printf("Error: unsupported target '%s'. Only 'claude-code' is currently supported.\n", target)
			return
		}
		teamConfig, err := getTeamConfig(teamName)
		if err != nil {
			fmt.Printf("Error getting team config: %s\n", err)
			return
		}
		exporter := &transporter.ClaudeCodeSubAgentExporter{
			TeamConfig: teamConfig,
			CodingAgent: &config.ClaudeCodingAgent{},
			FileUtils: &transporter.FileUtils{},
		}
		if projectDir == "" {
			fmt.Printf("Exporting subagent (%s) config file to user directory\n", subagentName)
			exporter.Export(subagentName)
		} else {
			fmt.Printf("Exporting subagent (%s) config file to project directory %s\n", subagentName, projectDir)
			exporter.ExportToProject(subagentName, projectDir)
		}
	},
}

func init() {
	// Export all command flags
	exportAllCmd.Flags().StringP("team", "t", "", "Team name")
	exportAllCmd.Flags().StringP("project-dir", "p", "", "Project directory")
	exportAllCmd.MarkFlagRequired("team")
	exportAllCmd.Flags().StringP("target", "c", "claude-code", "Coding agent target (claude-code etc.)")
	// Export subagent command flags
	exportSubagentCmd.Flags().StringP("name", "n", "", "Subagent name")
	exportSubagentCmd.MarkFlagRequired("name")
	exportSubagentCmd.Flags().StringP("team", "t", "", "Team name")
	exportSubagentCmd.MarkFlagRequired("team")
	exportSubagentCmd.Flags().StringP("project-dir", "p", "", "Project directory")
	exportSubagentCmd.Flags().StringP("target", "c", "claude-code", "Coding agent target (claude-code etc.)")
	// Add sub-commands
	exportCmd.AddCommand(exportAllCmd)
	exportCmd.AddCommand(exportSubagentCmd)
}
