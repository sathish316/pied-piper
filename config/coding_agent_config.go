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

type ClaudeCodingAgent struct {
}

func (c *ClaudeCodingAgent) GetName() string {
	return "claude-code"
}

func (c *ClaudeCodingAgent) GetUserSubagentConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".claude", "agents")
}

func (c *ClaudeCodingAgent) GetProjectSubagentConfigDir(projectDir string) string {
	return filepath.Join(projectDir, ".claude", "agents")
}