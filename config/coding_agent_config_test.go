package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClaudeCodingAgent(t *testing.T) {
	claudeCodingAgent := &ClaudeCodingAgent{}

	assert.Equal(t, "claude-code", claudeCodingAgent.GetName())
	assert.Equal(t, filepath.Join(os.Getenv("HOME"), ".claude", "agents"), claudeCodingAgent.GetUserSubagentConfigDir())

	projectDir := "/Users/john/code/projects/yet_another_b2b_saas_project"
	projectAgentsDir := fmt.Sprintf("%s/.claude/agents", projectDir)
	assert.Equal(t, projectAgentsDir, claudeCodingAgent.GetProjectSubagentConfigDir(projectDir))
}