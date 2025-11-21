package config

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v3"
)

//go:embed config.sample.yml
var sampleConfigContent []byte

const DEFAULT_CONFIG_DIR = ".pied-piper"
const DEFAULT_CONFIG_FILE = "config.yml"

type Config struct {
	Path string
	File string
	Data map[string]interface{}
}

func (c *Config) GetFilePath() string {
	return filepath.Join(c.Path, c.File)
}

func (c *Config) Init() error {
	os.MkdirAll(c.Path, 0755)
	fmt.Println("Initializing config file at ", c.GetFilePath())

	err := os.WriteFile(c.GetFilePath(), sampleConfigContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func (c *Config) Load() (map[string]interface{}, error) {
	data, err := os.ReadFile(c.GetFilePath())
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var result map[string]interface{}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	c.Data = result
	return result, nil
}

func (c *Config) PrettyPrint() (string, error) {
	output, err := yaml.Marshal(c.Data)
	if err != nil {
		return "", fmt.Errorf("error marshaling YAML: %w", err)
	}

	fmt.Println(string(output))
	return string(output), nil
}