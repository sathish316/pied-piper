package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupSubagentTest(t *testing.T) (*TeamConfigYamlHandler, string) {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "config-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	t.Cleanup(func() {
		os.RemoveAll(tmpDir)
	})

	configPath := TeamConfigPath{
		Path: tmpDir,
		File: DEFAULT_CONFIG_FILE,
	}

	teamConfigHandler := TeamConfigYamlHandler{
		ConfigPath: configPath,
	}

	return &teamConfigHandler, tmpDir
}

func TestSubagentConfigHandler_List(t *testing.T) {
	configHandler, _ := setupSubagentTest(t)
	err := configHandler.Init()
	assert.NoError(t, err)
	teamConfig, err := configHandler.Load()
	assert.NoError(t, err)
	subagentConfigHandler := SubagentConfigYamlHandler{
		Config: teamConfig,
	}

	subagents, err := subagentConfigHandler.List("pied-piper")

	assert.NoError(t, err)
	assert.Len(t, subagents, 6)
}

func TestSubagentConfigHandler_Show(t *testing.T) {
	configHandler, _ := setupSubagentTest(t)
	err := configHandler.Init()
	assert.NoError(t, err)
	teamConfig, err := configHandler.Load()
	assert.NoError(t, err)
	subagentConfigHandler := SubagentConfigYamlHandler{
		Config: teamConfig,
	}

	subagent, err := subagentConfigHandler.Show("pied-piper", "software-engineer")

	assert.NoError(t, err)
	assert.Equal(t, "software-engineer", subagent.Role)
	assert.Equal(t, "Gilfoyle", subagent.Nickname)
}
