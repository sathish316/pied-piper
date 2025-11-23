package config

import (
	"os"
	"path/filepath"
)

type CodingAgent interface {
	GetName() string
	GetUserSubagentConfigDir() string
	GetProjectSubagentConfigDir(projectDir string) string
}

const CLAUDE_CONFIG_DIR = ".claude"
const CLAUDE_AGENTS_DIR = "agents"

type ClaudeCodingAgent struct {
}

func (c *ClaudeCodingAgent) GetName() string {
	return string(ClaudeCode)
}

func (c *ClaudeCodingAgent) GetUserSubagentConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), CLAUDE_CONFIG_DIR, CLAUDE_AGENTS_DIR)
}

func (c *ClaudeCodingAgent) GetProjectSubagentConfigDir(projectDir string) string {
	return filepath.Join(projectDir, CLAUDE_CONFIG_DIR, CLAUDE_AGENTS_DIR)
}

type CodingAgentConfig struct {
	Target Target
	TargetDir string
	TargetDirType TargetDirType
}

func (c *CodingAgentConfig) ToString() string {
	return "Agent: " + string(c.Target) + ", TargetDir: " + c.TargetDir + ", TargetDirType: " + string(c.TargetDirType)
}

type Target string

const (
	ClaudeCode Target = "claude-code"
)

type TargetDirType string

const (
	TargetDirTypeUser TargetDirType = "user"
	TargetDirTypeProject TargetDirType = "project"
)