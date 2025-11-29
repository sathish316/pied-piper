package config

import (
	"fmt"
)

type CodingAgent interface {
	GetName() string
	GetUserSubagentConfigDir() string
	GetProjectSubagentConfigDir(projectDir string) string
	GetUserSubagentConfigFilePath(subagentName string) string
	GetProjectSubagentConfigFilePath(projectDir string, subagentName string) string
}

type CodingAgentConfig struct {
	Target Target
	TargetDir string
	TargetDirType TargetDirType
}

func (c *CodingAgentConfig) ToString() string {
	return fmt.Sprintf("%s (TargetDir(%s): %s)", string(c.Target), string(c.TargetDirType), c.TargetDir)
}

type Target string

const (
	ClaudeCode Target = "claude-code"
	Rovodev Target = "rovodev"
)

type TargetDirType string

const (
	TargetDirTypeUser TargetDirType = "user"
	TargetDirTypeProject TargetDirType = "project"
)