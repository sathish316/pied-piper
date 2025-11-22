package team

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestConfigHandler_Add(t *testing.T) {
	input := 5
	expected := 10

	result := add(input, input)

	assert.Equal(t, expected, result)
}

func setup(t *testing.T) (TeamConfigHandler, *Team, string) {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "config-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	t.Cleanup(func() {
		os.RemoveAll(tmpDir)
	})

	config := &TeamConfig{
		Path: tmpDir,
		File: DEFAULT_CONFIG_FILE,
	}

	team := &Team{
		TeamConfig: config,
	}

	teamConfigHandler := TeamConfigYamlHandler{
		Team: team,
	}
	return &teamConfigHandler, team, tmpDir
}

func TestConfigHandler_Init(t *testing.T) {
	configHandler, team, tmpDir := setup(t)
	err := configHandler.Init()
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(tmpDir, DEFAULT_CONFIG_FILE), team.GetConfigFilePath())
	assert.FileExists(t, team.GetConfigFilePath())
}

func TestConfigHandler_Load(t *testing.T) {
	configHandler, _, _ := setup(t)
	configHandler.Init()

	data, err := configHandler.Load()
	assert.NoError(t, err)

	assert.Equal(t, "pied-piper", data["name"])
}

func TestConfigHandler_PrettyPrint(t *testing.T) {
	configHandler, _, _ := setup(t)
	configHandler.Init()

	configHandler.Load()
	configStr, err := configHandler.PrettyPrint()

	assert.NoError(t, err)
	assert.Contains(t, configStr, `name: pied-piper`)
}

func add(a, b int) int {
	return a + b
}
