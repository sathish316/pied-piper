package config

import (
	_ "embed"
	"fmt"
	"os"
	"bytes"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

const DEFAULT_CONFIG_DIR = ".pied-piper"
const DEFAULT_CONFIG_FILE = "config.yml"

//go:embed config.sample.yml
var sampleConfigContent []byte

//go:embed config.blank.yml
var blankConfigContent []byte

// TeamConfigData represents the structured YAML data for team configuration
type TeamConfig struct {
	Name        string              `yaml:"name"`
	Description string        		`yaml:"description"`
	SubAgents   []SubagentConfig 	`yaml:"subagents"`
	ConfigPath	TeamConfigPath		`yaml:"-"`
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

func GetTeamConfigDir(teamName string) string {
	return filepath.Join(os.Getenv("HOME"), DEFAULT_CONFIG_DIR, teamName)
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
	var configContent []byte
	if c.Config != nil && c.Config.Name == "pied-piper" {
		configContent = sampleConfigContent
	} else {
		configContent = blankConfigContent
	}
	err = os.WriteFile(c.ConfigPath.GetConfigFilePath(), configContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	// Load the config so we can access team name and other data
	_, err = c.Load()
	if err != nil {
		return fmt.Errorf("error loading config after init: %w", err)
	}

	// Ensure subagents dir exists
	subagentsDir := filepath.Join(c.ConfigPath.Path, "subagents")
	fmt.Println("Initializing subagents dir at ", subagentsDir)
	err = os.MkdirAll(subagentsDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating subagents directory: %w", err)
	}
	// Ensure templates dir exists and initialize required templates
	templatesDir := filepath.Join(c.ConfigPath.Path, "templates")
	fmt.Println("Initializing templates dir at ", templatesDir)
	err = os.MkdirAll(templatesDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating templates directory: %w", err)
	}
	// Initialize subagent templates
	templatesConfig := TemplatesConfig{
		TeamName: c.Config.Name,
	}
	allCodingAgents := []Target{ClaudeCode, Rovodev}
	for _, codingAgent := range allCodingAgents {
		subagentTemplatePath := templatesConfig.GetSubagentTemplatePath(c.ConfigPath.Path, codingAgent)
		subagentTemplateContent, err := templatesConfig.GetSubagentTemplateContent(codingAgent)
		if err != nil {
			return fmt.Errorf("error getting subagent template for %s: %w", codingAgent, err)
		}
		err = os.WriteFile(subagentTemplatePath, subagentTemplateContent, 0644)
		fmt.Printf("Initializing subagent template for %s at %s\n", codingAgent, subagentTemplatePath)
		if err != nil {
			return fmt.Errorf("error writing subagent template for %s: %w", codingAgent, err)
		}
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

func (c *TeamConfigYamlHandler) Save() error {
    // Create a new encoder with custom formatting
    var buf bytes.Buffer
    encoder := yaml.NewEncoder(&buf)
    encoder.SetIndent(2) // Set indentation to 2 spaces

    // Encode the config
    err := encoder.Encode(c.Config)
    if err != nil {
        return fmt.Errorf("error marshalling team config: %w", err)
    }
    encoder.Close()

    // Write to file
    err = os.WriteFile(c.ConfigPath.GetConfigFilePath(), buf.Bytes(), 0644)
    if err != nil {
        return fmt.Errorf("error writing team config file: %w", err)
    }

    return nil
}
func (c *TeamConfigYamlHandler) PrettyPrint() (string, error) {
	config := c.Config
	output, err := yaml.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("error showing team config: %w", err)
	}
	return string(output), nil
}
