package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamConfigHandler_Add(t *testing.T) {
	input := 5
	expected := 10

	result := add(input, input)

	assert.Equal(t, expected, result)
}

func setup(t *testing.T) (*TeamConfigYamlHandler, string) {
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
		Config:     &TeamConfig{Name: "pied-piper"}, // Use pied-piper for sample config
	}
	return &teamConfigHandler, tmpDir
}

func TestTeamConfigHandler_Init(t *testing.T) {
	configHandler, tmpDir := setup(t)
	err := configHandler.Init()
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(tmpDir, DEFAULT_CONFIG_FILE), configHandler.ConfigPath.GetConfigFilePath())
	assert.FileExists(t, configHandler.ConfigPath.GetConfigFilePath())
}

func TestTeamConfigHandler_Load(t *testing.T) {
	configHandler, _ := setup(t)
	configHandler.Init()

	data, err := configHandler.Load()
	assert.NoError(t, err)

	assert.Equal(t, "pied-piper", data.Name)
	assert.Equal(t, "A team of AI SubAgents for SDLC workflow", data.Description)
	assert.Len(t, data.SubAgents, 7)
}

func TestTeamConfigHandler_PrettyPrint(t *testing.T) {
	configHandler, _ := setup(t)
	configHandler.Init()

	configHandler.Load()
	configStr, err := configHandler.PrettyPrint()

	assert.NoError(t, err)
	assert.Contains(t, configStr, `name: pied-piper`)
}

func add(a, b int) int {
	return a + b
}
