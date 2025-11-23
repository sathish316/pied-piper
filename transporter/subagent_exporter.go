package transporter

import (
	"fmt"
	"github.com/sathish316/pied-piper/config"
)

type SubAgentExporter interface {
	ExportAll() error
	Export(subagentName string) error
	ExportAllToProject(projectDir string) error
	ExportToProject(subagentName string, projectDir string) error
}

type ClaudeCodeSubAgentExporter struct {
	TeamConfig *config.TeamConfig
	CodingAgent config.CodingAgent
	*FileUtils
}

func (e *ClaudeCodeSubAgentExporter) ExportAll() error {
	// Export all subagent config files to user directory
	subagents := e.TeamConfig.SubAgents
	for _, subagent := range subagents {
		err := e.Export(subagent.Role)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *ClaudeCodeSubAgentExporter) Export(subagentName string) error {
	// Export single subagent config file to user directory
	fmt.Printf("Exporting subagent (%s) to %s user directory\n", subagentName, e.CodingAgent.GetName())
	srcFile := config.GetSubagentSpecFilePath(e.TeamConfig, subagentName)
	destDir := e.CodingAgent.GetUserSubagentConfigDir()
	return e.CopyFile(srcFile, destDir)
}

func (e *ClaudeCodeSubAgentExporter) ExportAllToProject(projectDir string) error {
	// Export all subagent config files to project directory
	subagents := e.TeamConfig.SubAgents
	for _, subagent := range subagents {
		err := e.ExportToProject(subagent.Role, projectDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *ClaudeCodeSubAgentExporter) ExportToProject(subagentName string, projectDir string) error {
	// Export single subagent config file to project directory
	fmt.Printf("Exporting subagent (%s) to %s project directory\n", subagentName, e.CodingAgent.GetName())
	srcFile := config.GetSubagentSpecFilePath(e.TeamConfig, subagentName)
	destDir := e.CodingAgent.GetProjectSubagentConfigDir(projectDir)
	return e.CopyFile(srcFile, destDir)
}
