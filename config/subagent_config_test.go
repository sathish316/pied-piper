package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"path/filepath"
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

func setupSubagentSpecTest(t *testing.T, teamConfigDir string) (string) {
	t.Helper()
	subagentSpecConfigPath := filepath.Join(teamConfigDir, "subagents", "software-engineer.yml")
	subagentSpecConfigContent := `
role: "software-engineer"
description: ""
nickname: "Gilfoyle"
task_labels:
  incoming:
  - "@ready-for-lld"
  - "@ready-for-dev"
  - "@code-review-rejected"
  outgoing:
  - "@ready-for-code-review"
wiki_labels:
  incoming:
  - "HLD_<TASK_ID>.md"
  - "LLD_<TASK_ID>.md"
  outgoing:
  - "LLD_<TASK_ID>.md"`
	err := os.WriteFile(subagentSpecConfigPath, []byte(subagentSpecConfigContent), 0644)
	fmt.Println("Test Setup: Writing subagent spec config file to ", subagentSpecConfigPath)
	if err != nil {
		t.Fatalf("Failed to create subagent spec config file: %v", err)
	}
	return subagentSpecConfigPath
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
	assert.Len(t, subagents, 7)
	assert.Equal(t, mapRoles(subagents), []string{"microsprint-orchestrator", "product-manager", "architect", "software-engineer", "code-reviewer", "code-validator", "build-engineer"})
}

func mapRoles(subagents []SubagentConfig) []string {
	roles := make([]string, 0)
	for _, subagent := range subagents {
		roles = append(roles, subagent.Role)
	}
	return roles
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
	assert.NotEmpty(t, subagent.TaskLabels.Incoming)
	assert.NotEmpty(t, subagent.TaskLabels.Outgoing)
}

func TestSubagentConfigHandler_GetSpec(t *testing.T) {
	configHandler, tmpDir := setupSubagentTest(t)
	err := configHandler.Init()
	assert.NoError(t, err)
	teamConfig, err := configHandler.Load()
	assert.NoError(t, err)
	setupSubagentSpecTest(t, tmpDir)
	subagentConfigHandler := SubagentConfigYamlHandler{
		Config: teamConfig,
	}

	subagentSpecConfig, err := subagentConfigHandler.GetSpec("pied-piper", "software-engineer")

	assert.NoError(t, err)
	assert.Equal(t, "software-engineer", subagentSpecConfig.Role)
	assert.Equal(t, "Gilfoyle", subagentSpecConfig.Nickname)
	assert.NotEmpty(t, subagentSpecConfig.TaskLabels.Incoming)
	assert.NotEmpty(t, subagentSpecConfig.TaskLabels.Outgoing)
}

func TestSubagentConfigHandler_UpdateSpec(t *testing.T) {
	configHandler, tmpDir := setupSubagentTest(t)
	err := configHandler.Init()
	assert.NoError(t, err)
	teamConfig, err := configHandler.Load()
	assert.NoError(t, err)
	setupSubagentSpecTest(t, tmpDir)
	subagentConfigHandler := SubagentConfigYamlHandler{
		Config: teamConfig,
	}

	architectSpec := SubagentSpecConfig{
		Role: "architect",
		Nickname: "Richard",
		TaskLabels: TaskLabelsConfig{
			Incoming: []string{"@ready-for-hld"},
			Outgoing: []string{"@ready-for-lld"},
		},
	}
	subagentSpecConfigPath, err := subagentConfigHandler.UpdateSpec("pied-piper", "architect", &architectSpec)

	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(teamConfig.ConfigPath.Path, "subagents", "architect.yml"), subagentSpecConfigPath)

	subagentSpecConfig, err := subagentConfigHandler.GetSpec("pied-piper", "architect")
	assert.NoError(t, err)
	assert.Equal(t, "architect", subagentSpecConfig.Role)
	assert.Equal(t, "Richard", subagentSpecConfig.Nickname)
	assert.NotEmpty(t, subagentSpecConfig.TaskLabels.Incoming)
	assert.NotEmpty(t, subagentSpecConfig.TaskLabels.Outgoing)
}
