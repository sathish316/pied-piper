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
	os.MkdirAll(c.ConfigPath.Path, 0755)
	fmt.Println("Initializing config file at ", c.ConfigPath.GetConfigFilePath())

	err := os.WriteFile(c.ConfigPath.GetConfigFilePath(), sampleConfigContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func (c *TeamConfigYamlHandler) Load() (*TeamConfig, error) {
	// Read YAML file
	yamlFile, err := os.ReadFile(c.ConfigPath.GetConfigFilePath())
	if err != nil {
		return nil, fmt.Errorf("Error reading team config file: %w", err)
	}

	// Initialize struct
	var teamConfig TeamConfig

	// Unmarshal YAML data into the struct
	err = yaml.Unmarshal(yamlFile, &teamConfig)
	if err != nil {
		return nil, fmt.Errorf("Error parsing config file: %w", err)
	}
	c.Config = &teamConfig
	return &teamConfig, nil
}

func (c *TeamConfigYamlHandler) PrettyPrint() (string, error) {
	config := c.Config
	output, err := yaml.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("Error showing team config: %w", err)
	}
	return string(output), nil
}
