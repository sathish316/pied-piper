package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type TaskLabelsConfig struct {
	Incoming []string `yaml:"incoming"`
	Outgoing []string `yaml:"outgoing"`
}

type WikiLabelsConfig struct {
	Incoming []string `yaml:"incoming"`
	Outgoing []string `yaml:"outgoing"`
}

type SubagentConfig struct {
	Role        string           `yaml:"role"`
	Description string           `yaml:"description"`
	Nickname    string           `yaml:"nickname"`
	TaskLabels  TaskLabelsConfig `yaml:"task_labels"`
	WikiLabels  WikiLabelsConfig `yaml:"wiki_labels"`
}

type SubagentSpecConfig struct {
	Role                string           `yaml:"role"`
	Description         string           `yaml:"description"`
	Nickname            string           `yaml:"nickname"`
	TaskLabels          TaskLabelsConfig `yaml:"task_labels"`
	WikiLabels          WikiLabelsConfig `yaml:"wiki_labels"`
	WorkflowDescription string           `yaml:"workflow_description"`
	RoleDescription     string           `yaml:"role_description"`
	Memory              string           `yaml:"memory"`
}

func (s *SubagentConfig) ToString() string {
	yamlStr, err := yaml.Marshal(s)
	if err != nil {
		return ""
	}
	return string(yamlStr)
}

type SubagentConfigHandler interface {
	List(teamName string) ([]SubagentConfig, error)
	Show(teamName string, subagentName string) (*SubagentConfig, error)
	GetSpec(teamName string, subagentName string) (*SubagentSpecConfig, error)
	UpdateSpec(teamName string, subagentName string, subagentSpec *SubagentSpecConfig) (string, error)
	UpdateSpecYaml(teamName string, subagentName string, yamlStr []byte) (string, error)
}

type SubagentConfigYamlHandler struct {
	Config *TeamConfig
}

func (c *SubagentConfigYamlHandler) List(teamName string) ([]SubagentConfig, error) {
	//FIXME: Make this work for multiple teams
	subagents := c.Config.SubAgents
	return subagents, nil
}

func (c *SubagentConfigYamlHandler) Show(teamName string, subagentName string) (*SubagentConfig, error) {
	//FIXME: Make this work for multiple teams
	subagents := c.Config.SubAgents
	for _, subagent := range subagents {
		if subagent.Role == subagentName || subagent.Nickname == subagentName {
			return &subagent, nil
		}
	}
	return nil, fmt.Errorf("subagent %s not found", subagentName)
}

func (c *SubagentConfigYamlHandler) GetSpec(teamName string, subagentName string) (*SubagentSpecConfig, error) {
	//FIXME: Make this work for multiple teams
	//FIXME: Make this work for multiple subagents with same role
	// Go to <team-config-dir>/subagents/<subagent-name>.yml
	// Read YAML file
	teamConfigPath := c.Config.ConfigPath.Path
	subagentSpecConfigPath := filepath.Join(teamConfigPath, "subagents", subagentName+".yml")
	yamlFile, err := os.ReadFile(subagentSpecConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error reading subagent spec config file: %w", err)
	}

	// Initialize struct
	var subagentSpecConfig SubagentSpecConfig

	// Unmarshal YAML data into the struct
	err = yaml.Unmarshal(yamlFile, &subagentSpecConfig)
	if err != nil {
		return nil, fmt.Errorf("error parsing subagent spec config file: %w", err)
	}
	return &subagentSpecConfig, nil
}

func (c *SubagentConfigYamlHandler) UpdateSpec(teamName string, subagentName string, subagentSpec *SubagentSpecConfig) (string, error) {
	//FIXME: Make this work for multiple teams
	//FIXME: Make this work for multiple subagents with same role
	// Marshal subagentSpec to YAML
	yamlStr, err := yaml.Marshal(subagentSpec)
	if err != nil {
		return "", fmt.Errorf("error marshalling subagent spec config: %w", err)
	}
	return c.UpdateSpecYaml(teamName, subagentName, yamlStr)
}

func (c *SubagentConfigYamlHandler) UpdateSpecYaml(teamName string, subagentName string, yamlStr []byte) (string, error) {
	//FIXME: Make this work for multiple teams
	//FIXME: Make this work for multiple subagents with same role
	// Go to <team-config-dir>/subagents/<subagent-name>.yml
	// Write YAML file
	teamConfigPath := c.Config.ConfigPath.Path
	subagentSpecConfigPath := filepath.Join(teamConfigPath, "subagents", subagentName+".yml")
	fmt.Println("Updating subagent-spec config file at ", subagentSpecConfigPath)
	err := os.WriteFile(subagentSpecConfigPath, yamlStr, 0644)
	if err != nil {
		return "", fmt.Errorf("error writing subagent spec config file: %w", err)
	}
	return subagentSpecConfigPath, nil
}
