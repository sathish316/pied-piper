package config

import (
	_ "embed"
	"fmt"
	"path/filepath"
)

//go:embed subagent_template_claude-code.md
var subagentTemplateClaudeCodeContent []byte

const TEMPLATES_DIR = "templates"

type TemplatesConfig struct {
	TeamName string
}

func (t *TemplatesConfig) GetSubagentTemplatePath(teamConfigPath string, codingAgentTarget Target) string {
	return filepath.Join(
		teamConfigPath,
		TEMPLATES_DIR,
		fmt.Sprintf("subagent_template_%s.md", string(codingAgentTarget)),
	)
}

func (t *TemplatesConfig) GetSubagentTemplateContent(codingAgentTarget Target) ([]byte, error) {
	if codingAgentTarget == ClaudeCode {
		return subagentTemplateClaudeCodeContent, nil
	}
	return nil, fmt.Errorf("subagent template for %s is not found", codingAgentTarget)
}
