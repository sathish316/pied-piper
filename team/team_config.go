package team

import (
	_ "embed"
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

//go:embed config.sample.yml
var sampleConfigContent []byte

type TeamConfig struct {
	Path string
	File string
	Data map[string]interface{}
}

type TeamConfigHandler interface {
	Init() error
	Load() (map[string]interface{}, error)
	PrettyPrint() (string, error)
}

type TeamConfigYamlHandler struct {
	Team *Team
}

func (c *TeamConfigYamlHandler) Init() error {
	os.MkdirAll(c.Team.GetConfigDir(), 0755)
	fmt.Println("Initializing config file at ", c.Team.GetConfigFilePath())

	err := os.WriteFile(c.Team.GetConfigFilePath(), sampleConfigContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func (c *TeamConfigYamlHandler) Load() (map[string]interface{}, error) {
	data, err := os.ReadFile(c.Team.GetConfigFilePath())
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var result map[string]interface{}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	c.Team.SetConfigData(result)
	return result, nil
}

func (c *TeamConfigYamlHandler) PrettyPrint() (string, error) {
	output, err := yaml.Marshal(c.Team.GetConfigData())
	if err != nil {
		return "", fmt.Errorf("error marshaling YAML: %w", err)
	}

	fmt.Println(string(output))
	return string(output), nil
}