package transporter

import "github.com/sathish316/pied-piper/config"

type SubAgentImporter interface {
	ImportAll(teamConfig *config.TeamConfig) error
	Import(teamConfig *config.TeamConfig, subagentName string) error
}

type ClaudeCodeSubAgentImporter struct {
	TeamConfig *config.TeamConfig
}

func (i *ClaudeCodeSubAgentImporter) ImportAll() error {
	// Import all subagent config files from user directory
	return nil
}

func (i *ClaudeCodeSubAgentImporter) Import(subagentName string) error {
	// Import single subagent config file from user directory
	return nil
}

func (i *ClaudeCodeSubAgentImporter) ImportAllFromProject(projectDir string) error {
	// Import all subagent config files from project directory
	return nil
}

func (i *ClaudeCodeSubAgentImporter) ImportFromProject(subagentName string, projectDir string) error {
	// Import single subagent config file from project directory
	return nil
}