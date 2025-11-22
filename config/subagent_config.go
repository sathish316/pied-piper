package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type SubagentConfig struct {
	Role		string				`yaml:"role"`
	Description string              `yaml:"description"`
	Nickname	string				`yaml:"nickname"`
	Prompt      string              `yaml:"prompt"`
	Tools       []string            `yaml:"tools"`
	TaskLabels  map[string][]string `yaml:"task_labels"`
	WikiLabels  map[string][]string `yaml:"wiki_labels"`
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
	return nil, fmt.Errorf("Subagent %s not found", subagentName)
}
