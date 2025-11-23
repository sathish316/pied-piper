package config

import (
	_ "embed"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"path/filepath"
)

const DEFAULT_CONFIG_DIR = ".pied-piper"
const DEFAULT_CONFIG_FILE = "config.yml"

//go:embed config.sample.yml
var sampleConfigContent []byte

// TeamConfigData represents the structured YAML data for team configuration
type TeamConfig struct {
	Name        string              `yaml:"name"`
	Description string        		`yaml:"description"`
	SubAgents   []SubagentConfig 	`yaml:"subagents"`
	ConfigPath	TeamConfigPath
}

func (t *TeamConfig) FindSubagentByRole(role string) (*SubagentConfig, error) {
	for _, subagent := range t.SubAgents {
		if subagent.Role == role {
			return &subagent, nil
		}
	}
	return nil, fmt.Errorf("subagent %s not found", role)
}

type TeamConfigPath struct {
	Path string
	File string
}

func (t *TeamConfigPath) GetConfigFilePath() string {
	return filepath.Join(t.Path, t.File)
}

type TeamConfigHandler interface {
	Init() error
	Load() (*TeamConfigPath, error)
	PrettyPrint() (string, error)
}

type TeamConfigYamlHandler struct {
	ConfigPath TeamConfigPath
	Config *TeamConfig
}

func (c *TeamConfigYamlHandler) Init() error {
	fmt.Println("Initializing config file at ", c.ConfigPath.GetConfigFilePath())
	// Ensure team config dir exists
	err := os.MkdirAll(c.ConfigPath.Path, 0755)
	if err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	// Ensure team config file exists
	err = os.WriteFile(c.ConfigPath.GetConfigFilePath(), sampleConfigContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	// Ensure subagents dir exists
	err = os.MkdirAll(filepath.Join(c.ConfigPath.Path, "subagents"), 0755)
	if err != nil {
		return fmt.Errorf("error creating subagents directory: %w", err)
	}
	return nil
}

func (c *TeamConfigYamlHandler) Load() (*TeamConfig, error) {
	// Read YAML file
	yamlFile, err := os.ReadFile(c.ConfigPath.GetConfigFilePath())
	if err != nil {
		return nil, fmt.Errorf("error reading team config file: %w", err)
	}

	// Initialize struct
	var teamConfig TeamConfig

	// Unmarshal YAML data into the struct
	err = yaml.Unmarshal(yamlFile, &teamConfig)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}
	c.Config = &teamConfig
	teamConfig.ConfigPath = c.ConfigPath
	return &teamConfig, nil
}

func (c *TeamConfigYamlHandler) PrettyPrint() (string, error) {
	config := c.Config
	output, err := yaml.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("error showing team config: %w", err)
	}
	return string(output), nil
}
