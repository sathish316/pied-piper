package cli

import (
	"fmt"

	"github.com/spf13/cobra"
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
		projectDir, _ := cmd.Flags().GetString("project-dir")
		if projectDir == "" {
			fmt.Println("Exporting all subagent config files to user directory")
		} else {
			fmt.Printf("Exporting all subagent config files to project directory %s\n", projectDir)
		}
	},
}

var exportSubagentCmd = &cobra.Command{
	Use:   "subagent",
	Short: "Export subagent config file",
	Long:  `Export subagent config file`,
	Run: func(cmd *cobra.Command, args []string) {
		subagentName, _ := cmd.Flags().GetString("name")
		projectDir, _ := cmd.Flags().GetString("project-dir")
		if projectDir == "" {
			fmt.Printf("Exporting subagent (%s) config file to user directory\n", subagentName)
		} else {
			fmt.Printf("Exporting subagent (%s) config file to project directory %s\n", subagentName, projectDir)
		}
	},
}

func init() {
	// Export all command flags
	exportAllCmd.Flags().StringP("project-dir", "p", "", "Project directory")
	// Export subagent command flags
	exportSubagentCmd.Flags().StringP("name", "n", "", "Subagent name")
	exportSubagentCmd.MarkFlagRequired("name")
	exportSubagentCmd.Flags().StringP("project-dir", "p", "", "Project directory")
	// Add sub-commands
	exportCmd.AddCommand(exportAllCmd)
	exportCmd.AddCommand(exportSubagentCmd)
}
