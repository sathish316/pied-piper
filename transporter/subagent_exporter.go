package transporter

import (
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
}

func (e *ClaudeCodeSubAgentExporter) ExportAll() error {
	// Export all subagent config files to user directory
	return nil
}

func (e *ClaudeCodeSubAgentExporter) Export(subagentName string) error {
	// Export single subagent config file to user directory
	return nil
}

func (e *ClaudeCodeSubAgentExporter) ExportAllToProject(projectDir string) error {
	// Export all subagent config files to project directory
	return nil
}

func (e *ClaudeCodeSubAgentExporter) ExportToProject(subagentName string, projectDir string) error {
	// Export single subagent config file to project directory
	return nil
}