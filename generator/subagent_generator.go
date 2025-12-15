package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/sathish316/pied-piper/config"
	"github.com/sathish316/pied-piper/subagent"
)

type SubAgentGenerator interface {
	GenerateSubagentSpec(role string) (subagent.SubagentSpec, error)

	GenerateWorkflowSpec(subagentSpec subagent.SubagentSpec) (subagent.SubagentSpec, error)

	GenerateRoleSpec(subagentSpec subagent.SubagentSpec) (subagent.SubagentSpec, error)

	GenerateSubagentYaml(subagentSpec subagent.SubagentSpec) (string, error)

	GenerateSubagentSpecForCodingAgent(subagentSpec subagent.SubagentSpec, codingAgentConfig *config.CodingAgentConfig) (string, error)
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
		Role:                             subagentConfig.Role,
		Nickname:                         subagentConfig.Nickname,
		Description:                      subagentConfig.Description,
		TaskLabels:                       subagentConfig.TaskLabels,
		WikiLabels:                       subagentConfig.WikiLabels,
		GeneratedTaskWorkflowDescription: "",
		GeneratedWikiWorkflowDescription: "",
		RoleDescription:                  "",
		Memory:                           "",
		Model:                            subagentConfig.Model,
		RouterModel:                      subagentConfig.RouterModel,
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
	subagentYmlFilePath, err := subagentConfigHandler.UpdateSpec(g.TeamConfig.Name, subagentSpec.Role, subagentSpec.ToConfig())
	if err != nil {
		return "", err
	}
	return subagentYmlFilePath, nil
}

func (g *SDLCSubAgentGenerator) GenerateSubagentSpecFileForCodingAgent(subagentConfig *config.SubagentConfig, codingAgentConfig *config.CodingAgentConfig) (string, error) {
	fmt.Printf("Generating subagent spec (%s) for coding agent: %s\n", subagentConfig.Role, codingAgentConfig.Target)
	// Get pied-piper config dir
	subagentConfigHandler := config.SubagentConfigYamlHandler{
		Config: g.TeamConfig,
	}
	// Convert SubagentConfig to SubagentSpec
	subagentSpec, err := g.GenerateSubagentSpec(subagentConfig.Role)
	if err != nil {
		return "", err
	}
	// Generate subagent yaml specific to coding agent
	subagentMD, err := g.generateSubagentMDSpec(subagentSpec, codingAgentConfig)
	if err != nil {
		return "", err
	}
	//TODO: use codingAgent or LLM to generate other parts of the subagent yaml
	subagentMDFilePath, err := subagentConfigHandler.UpdateSpecMD(g.TeamConfig.Name, subagentSpec.Role, []byte(subagentMD))
	if err != nil {
		return "", err
	}
	return subagentMDFilePath, nil
}

func (g *SDLCSubAgentGenerator) generateSubagentMDSpec(subagentSpec *subagent.SubagentSpec, codingAgentConfig *config.CodingAgentConfig) (string, error) {
	// Load subagent template for Coding Agent
	templatesConfig := config.TemplatesConfig{
		TeamName: g.TeamConfig.Name,
	}
	subagentTemplatePath := templatesConfig.GetSubagentTemplatePath(g.TeamConfig.ConfigPath.Path, codingAgentConfig.Target)
	subagentTemplateContent, err := os.ReadFile(subagentTemplatePath)
	if err != nil {
		return "", err
	}
	// Create Text Template - https://pkg.go.dev/text/template
	subagentTemplate := template.New("subagentTemplate")
	subagentTemplate.Parse(string(subagentTemplateContent))
	var buf bytes.Buffer
	// Render Text Template
	err = subagentTemplate.Execute(&buf, subagentSpec)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
