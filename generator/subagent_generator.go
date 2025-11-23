package generator

import (
	"github.com/sathish316/pied-piper/subagent"
	"github.com/sathish316/pied-piper/config"
)

type SubAgentGenerator interface {
	GenerateSubagentSpec(role string) (subagent.SubagentSpec, error)

	GenerateWorkflowSpec(subagentSpec subagent.SubagentSpec) (subagent.SubagentSpec, error)

	GenerateRoleSpec(subagentSpec subagent.SubagentSpec) (subagent.SubagentSpec, error)

	GenerateSubagentYaml(subagentSpec subagent.SubagentSpec) (string, error)
}

type SDLCSubAgentGenerator struct {
	TeamConfig *config.TeamConfig
}

func (g *SDLCSubAgentGenerator) GenerateSubagentSpec(role string) (*subagent.SubagentSpec, error) {
	subagentConfig, err := g.TeamConfig.FindSubagentByRole(role)
	if err != nil {
		return nil, err
	}
	subagentSpec := subagent.SubagentSpec{
		Role: subagentConfig.Role,
		Nickname: subagentConfig.Nickname,
		Description: subagentConfig.Description,
		TaskLabels: subagentConfig.TaskLabels,
		WikiLabels: subagentConfig.WikiLabels,
		WorkflowDescription: "",
		RoleDescription: "",
		Memory: "",
	}
	return &subagentSpec, nil
}

func (g *SDLCSubAgentGenerator) GenerateWorkflowSpec(subagentSpec subagent.SubagentSpec) (subagent.SubagentSpec, error) {
	return subagent.SubagentSpec{}, nil
}

func (g *SDLCSubAgentGenerator) GenerateRoleSpec(subagentSpec subagent.SubagentSpec) (subagent.SubagentSpec, error) {
	return subagent.SubagentSpec{}, nil
}

func (g *SDLCSubAgentGenerator) GenerateSubagentYaml(subagentSpec *subagent.SubagentSpec) (string, error) {
	// Get pied-piper config dir
	subagentConfigHandler := config.SubagentConfigYamlHandler{
		Config: g.TeamConfig,
	}
	subagentConfigFilePath, err := subagentConfigHandler.UpdateSpec(g.TeamConfig.Name, subagentSpec.Role, subagentSpec.ToConfig())
	if err != nil {
		return "", err
	}
	return subagentConfigFilePath, nil
}

func (g *SDLCSubAgentGenerator) GenerateSubagentYamlForCodingAgent(subagentConfig *config.SubagentConfig, codingAgentConfig config.CodingAgentConfig) (string, error) {
	// Convert SubagentConfig to SubagentSpec
	subagentSpec, err := g.GenerateSubagentSpec(subagentConfig.Role)
	if err != nil {
		return "", err
	}
	//TODO: use codingAgent to generate subagent yaml specific to claude-code
	subagentYamlFilePath, err := g.GenerateSubagentYaml(subagentSpec)
	if err != nil {
		return "", err
	}
	return subagentYamlFilePath, nil
}