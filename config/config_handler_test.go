package config

import (
	"testing"
	"os"
	"path/filepath"
	"github.com/stretchr/testify/assert"
)

func TestConfigHandler_Add(t *testing.T) {
	input := 5
	expected := 10

	result := add(input, input)

	assert.Equal(t, expected, result)
}

func TestConfigHandler_Init(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "config-test-*")
	defer os.RemoveAll(tmpDir)
	config := Config{
		Path: tmpDir,
		File: DEFAULT_CONFIG_FILE,
	}
	config.Init()
	assert.Equal(t, config.GetFilePath(), filepath.Join(tmpDir, DEFAULT_CONFIG_FILE))
	assert.FileExists(t, config.GetFilePath())
}

func add(a, b int) int {
	return a + b
}