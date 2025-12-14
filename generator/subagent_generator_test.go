package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sathish316/pied-piper/config"
	"github.com/stretchr/testify/assert"
)

func setupSubagentGenTest(t *testing.T) (*SDLCSubAgentGenerator, *config.TeamConfig, string) {
	t.Helper()

	// Create team config in temp dir
	tmpDir, err := os.MkdirTemp("", "config-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	t.Cleanup(func() {
		os.RemoveAll(tmpDir)
	})

	configPath := config.TeamConfigPath{
		Path: tmpDir,
		File: config.DEFAULT_CONFIG_FILE,
	}

	// Load team config
	teamConfigHandler := config.TeamConfigYamlHandler{
		ConfigPath: configPath,
		Config:     &config.TeamConfig{Name: "pied-piper"}, // Use pied-piper for sample config
	}
	err = teamConfigHandler.Init()
	assert.NoError(t, err)
	teamConfig, err := teamConfigHandler.Load()
	assert.NoError(t, err)

	// Create subagent generator
	generator := &SDLCSubAgentGenerator{
		TeamConfig: teamConfig,
	}

	return generator, teamConfig, tmpDir
}
func TestSubAgentGenerator_GenerateSubagentSpec(t *testing.T) {
	generator, _, _ := setupSubagentGenTest(t)
	subagentSpec, err := generator.GenerateSubagentSpec("software-engineer")

	if assert.NoError(t, err) {
		assert.Equal(t, "software-engineer", subagentSpec.Role)
		assert.Equal(t, "Gilfoyle", subagentSpec.Nickname)
		// assert.Equal(t, "A software engineer who writes code", subagentSpec.Description)
	}
}

func TestSubAgentGenerator_GenerateSubagentSpecYamlFile(t *testing.T) {
	generator, teamConfig, _ := setupSubagentGenTest(t)
	assert.NotNil(t, teamConfig)
	subagentSpec, err := generator.GenerateSubagentSpec("software-engineer")
	assert.NoError(t, err)
	subagentSpecConfigPath, err := generator.GenerateSubagentYaml(subagentSpec)

	assert.NoError(t, err)

	// cat yaml content into memory and assert on the content
	yamlContent, err := os.ReadFile(subagentSpecConfigPath)
	assert.NoError(t, err)
	assert.Contains(t, string(yamlContent), "role: software-engineer")
	assert.Contains(t, string(yamlContent), "nickname: Gilfoyle")
	assert.Contains(t, string(yamlContent), "@ready-for-dev")
	assert.Contains(t, string(yamlContent), "@ready-for-code-review")
}

func TestSubAgentGenerator_GenerateSubagentSpecYamlFileForCodingTemplate_UsesTemplate(t *testing.T) {
	generator, teamConfig, tmpDir := setupSubagentGenTest(t)
	assert.NotNil(t, teamConfig)
	subagentConfig, err := teamConfig.FindSubagentByRole("software-engineer")
	assert.NoError(t, err)
	targetConfig := config.CodingAgentConfig{
		Target:    config.ClaudeCode,
		TargetDir: filepath.Join(tmpDir, config.CLAUDE_CONFIG_DIR, config.CLAUDE_AGENTS_DIR),
	}

	subagentSpecConfigPath, err := generator.GenerateSubagentSpecFileForCodingAgent(subagentConfig, &targetConfig)

	assert.NoError(t, err)
	// cat yaml content into memory and assert on the content
	yamlContent, err := os.ReadFile(subagentSpecConfigPath)
	assert.NoError(t, err)
	assert.Contains(t, string(yamlContent), "name: software-engineer")
	assert.Contains(t, string(yamlContent), "nickname: Gilfoyle")
	assert.Contains(t, string(yamlContent), "----TASK_WORKFLOW_DESCRIPTION STARTS----")
	assert.Contains(t, string(yamlContent), "----ROLE_DESCRIPTION STARTS----")
}

func TestSubAgentGenerator_GenerateSubagentSpecFileForCodingAgent_WithDifferentModels(t *testing.T) {
	// Setup: Create a temp dir and initialize team config
	tmpDir, err := os.MkdirTemp("", "config-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	t.Cleanup(func() {
		os.RemoveAll(tmpDir)
	})

	configPath := config.TeamConfigPath{
		Path: tmpDir,
		File: config.DEFAULT_CONFIG_FILE,
	}

	teamConfigHandler := config.TeamConfigYamlHandler{
		ConfigPath: configPath,
	}
	err = teamConfigHandler.Init()
	assert.NoError(t, err)
	teamConfig, err := teamConfigHandler.Load()
	assert.NoError(t, err)

	// Setup 3 subagents with different models
	teamConfig.SubAgents = []config.SubagentConfig{
		{
			Role:        "architect",
			Nickname:    "Richard",
			Description: "System architect",
			Model:       "opus",
			TaskLabels:  config.TaskLabelsConfig{Incoming: []string{"@ready-for-hld"}, Outgoing: []string{"@ready-for-lld"}},
		},
		{
			Role:        "software-engineer",
			Nickname:    "Gilfoyle",
			Description: "Software engineer",
			Model:       "sonnet",
			TaskLabels:  config.TaskLabelsConfig{Incoming: []string{"@ready-for-dev"}, Outgoing: []string{"@ready-for-code-review"}},
		},
		{
			Role:        "code-reviewer",
			Nickname:    "Dinesh",
			Description: "Code reviewer",
			Model:       "haiku",
			TaskLabels:  config.TaskLabelsConfig{Incoming: []string{"@ready-for-code-review"}, Outgoing: []string{"@code-review-done"}},
		},
	}

	generator := &SDLCSubAgentGenerator{
		TeamConfig: teamConfig,
	}

	targetConfig := config.CodingAgentConfig{
		Target:    config.ClaudeCode,
		TargetDir: filepath.Join(tmpDir, config.CLAUDE_CONFIG_DIR, config.CLAUDE_AGENTS_DIR),
	}

	// Generate spec files for all 3 subagents and verify each has the correct model
	testCases := []struct {
		role          string
		expectedModel string
	}{
		{"architect", "opus"},
		{"software-engineer", "sonnet"},
		{"code-reviewer", "haiku"},
	}

	for _, tc := range testCases {
		t.Run(tc.role, func(t *testing.T) {
			subagentConfig, err := teamConfig.FindSubagentByRole(tc.role)
			assert.NoError(t, err)

			subagentSpecConfigPath, err := generator.GenerateSubagentSpecFileForCodingAgent(subagentConfig, &targetConfig)
			assert.NoError(t, err)

			// Read the generated file and verify the model
			yamlContent, err := os.ReadFile(subagentSpecConfigPath)
			assert.NoError(t, err)
			assert.Contains(t, string(yamlContent), "model: "+tc.expectedModel,
				"Expected model %s for role %s, but got different model in content:\n%s",
				tc.expectedModel, tc.role, string(yamlContent))
		})
	}
}
