package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import subagent config files",
	Long:  `Import subagent config files`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var importAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Import all subagent config files",
	Long:  `Import all subagent config files`,
	Run: func(cmd *cobra.Command, args []string) {
		projectDir, _ := cmd.Flags().GetString("project-dir")
		if projectDir == "" {
			fmt.Println("Importing all subagent config files from user directory")
		} else {
			fmt.Printf("Importing all subagent config files from project directory %s\n", projectDir)
		}
	},
}

var importSubagentCmd = &cobra.Command{
	Use:   "subagent",
	Short: "Import subagent config file",
	Long:  `Import subagent config file`,
	Run: func(cmd *cobra.Command, args []string) {
		subagentName, _ := cmd.Flags().GetString("name")
		projectDir, _ := cmd.Flags().GetString("project-dir")
		if projectDir == "" {
			fmt.Printf("Importing subagent (%s) config file from user directory\n", subagentName)
		} else {
			fmt.Printf("Importing subagent (%s) config file from project directory %s\n", subagentName, projectDir)
		}
	},
}

func init() {
	// Import all command flags
	importAllCmd.Flags().StringP("project-dir", "p", "", "Project directory")
	// Import subagent command flags
	importSubagentCmd.Flags().StringP("name", "n", "", "Subagent name")
	importSubagentCmd.MarkFlagRequired("name")
	importSubagentCmd.Flags().StringP("project-dir", "p", "", "Project directory")
	// Add sub-commands
	importCmd.AddCommand(importAllCmd)
	importCmd.AddCommand(importSubagentCmd)
}
