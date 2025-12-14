package generator

import (
	"os"
	"testing"
	"github.com/sathish316/pied-piper/config"
	"github.com/stretchr/testify/assert"
	"path/filepath"
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

	assert.NoError(t, err)
	assert.Equal(t, "software-engineer", subagentSpec.Role)
	assert.Equal(t, "Gilfoyle", subagentSpec.Nickname)
	// assert.Equal(t, "A software engineer who writes code", subagentSpec.Description)
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
		Target: config.ClaudeCode,
		TargetDir: filepath.Join(tmpDir, config.CLAUDE_CONFIG_DIR, config.CLAUDE_AGENTS_DIR),
	}

	subagentSpecConfigPath, err := generator.GenerateSubagentSpecFileForCodingAgent(subagentConfig, &targetConfig)

	assert.NoError(t, err)
	// cat yaml content into memory and assert on the content
	yamlContent, err := os.ReadFile(subagentSpecConfigPath)
	assert.NoError(t, err)
	assert.Contains(t, string(yamlContent), "name: software-engineer")
	assert.Contains(t, string(yamlContent), "nickname: Gilfoyle")
	assert.Contains(t, string(yamlContent), "@ready-for-dev")
	assert.Contains(t, string(yamlContent), "@ready-for-code-review")
	assert.Contains(t, string(yamlContent), "----WORKFLOW_DESCRIPTION STARTS----")
	assert.Contains(t, string(yamlContent), "----ROLE_DESCRIPTION STARTS----")
}