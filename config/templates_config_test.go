package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplatesConfig_GetSubagentTemplatePath(t *testing.T) {
	teamConfig := &TeamConfig{
		ConfigPath: TeamConfigPath{
			Path: "/Users/john/.pied-piper/pied-piper",
		},
	}
	templatesConfig := &TemplatesConfig{
		TeamName: "pied-piper",
	}
	assert.Equal(t, "/Users/john/.pied-piper/pied-piper/templates/subagent_template_claude-code.md", templatesConfig.GetSubagentTemplatePath(teamConfig.ConfigPath.Path, ClaudeCode))
}
