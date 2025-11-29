package config

import (
	"os"
	"path/filepath"
)

const CLAUDE_CONFIG_DIR = ".claude"
const CLAUDE_AGENTS_DIR = "agents"

type ClaudeCodingAgent struct {
}

func (c ClaudeCodingAgent) GetName() string {
	return string(ClaudeCode)
}

func (c ClaudeCodingAgent) GetUserSubagentConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), CLAUDE_CONFIG_DIR, CLAUDE_AGENTS_DIR)
}

func (c ClaudeCodingAgent) GetProjectSubagentConfigDir(projectDir string) string {
	return filepath.Join(projectDir, CLAUDE_CONFIG_DIR, CLAUDE_AGENTS_DIR)
}

func (c ClaudeCodingAgent) GetUserSubagentConfigFilePath(subagentName string) string {
	return filepath.Join(c.GetUserSubagentConfigDir(), subagentName+".md")
}

func (c ClaudeCodingAgent) GetProjectSubagentConfigFilePath(projectDir string, subagentName string) string {
	return filepath.Join(c.GetProjectSubagentConfigDir(projectDir), subagentName+".md")
}
