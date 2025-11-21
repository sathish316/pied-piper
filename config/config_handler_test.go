package config

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

func setup(t *testing.T) (*Config, string) {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "config-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	t.Cleanup(func() {
		os.RemoveAll(tmpDir)
	})

	config := &Config{
		Path: tmpDir,
		File: DEFAULT_CONFIG_FILE,
	}

	return config, tmpDir
}

func TestConfigHandler_Init(t *testing.T) {
	config, tmpDir := setup(t)
	err := config.Init()
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(tmpDir, DEFAULT_CONFIG_FILE), config.GetFilePath())
	assert.FileExists(t, config.GetFilePath())
}

// func TestConfigHandler_Load(t *testing.T) {
// 	config, _ := setup(t)
// 	config.Init()

// 	data, err := config.Load()
// 	assert.NoError(t, err)

// 	assert.Equal(t, "pied-piper", data["name"])
// }

// func TestConfigHandler_PrettyPrint(t *testing.T) {
// 	config, _ := setup(t)
// 	configStr, err := config.PrettyPrint()

// 	assert.NoError(t, err)
// 	assert.Contains(t, configStr, `name: "pied-piper"`)
// }

func add(a, b int) int {
	return a + b
}
