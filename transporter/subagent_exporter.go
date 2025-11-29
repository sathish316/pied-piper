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

type SubAgentExporterImpl struct {
	TeamConfig *config.TeamConfig
	CodingAgent config.CodingAgent
	*FileUtils
}

func (e *SubAgentExporterImpl) ExportAll() error {
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

func (e *SubAgentExporterImpl) Export(subagentName string) error {
	// Export single subagent config file to user directory
	srcFile := config.GetSubagentSpecFilePath(e.TeamConfig, subagentName)
	destDir := e.CodingAgent.GetUserSubagentConfigFilePath(subagentName)
	fmt.Printf("Exporting subagent (%s) to %s user directory - %s\n", subagentName, e.CodingAgent.GetName(), destDir)
	return e.CopyFile(srcFile, destDir)
}

func (e *SubAgentExporterImpl) ExportAllToProject(projectDir string) error {
	// Export all subagent config files to project directory
	subagents := e.TeamConfig.SubAgents
	for _, subagent := range subagents {
		claude_code_exporter := &SubAgentExporterImpl{
			TeamConfig: e.TeamConfig,
			CodingAgent: e.CodingAgent,
			FileUtils: &FileUtils{},
		}
		exporter := claude_code_exporter
		err := exporter.ExportToProject(subagent.Role, projectDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *SubAgentExporterImpl) ExportToProject(subagentName string, projectDir string) error {
	// Export single subagent config file to project directory
	srcFile := config.GetSubagentSpecFilePath(e.TeamConfig, subagentName)
	destDir := e.CodingAgent.GetProjectSubagentConfigFilePath(projectDir, subagentName)
	fmt.Printf("Exporting subagent (%s) to %s project directory - %s\n", subagentName, e.CodingAgent.GetName(), projectDir)
	return e.CopyFile(srcFile, destDir)
}
