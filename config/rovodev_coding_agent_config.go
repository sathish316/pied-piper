package config

import (
	"os"
	"path/filepath"
)

const ROVODEV_CONFIG_DIR = ".rovodev"
const ROVODEV_AGENTS_DIR = "subagents"


type RovodevCodingAgent struct {
}

func (c RovodevCodingAgent) GetName() string {
	return string(Rovodev)
}

func (c RovodevCodingAgent) GetUserSubagentConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ROVODEV_CONFIG_DIR, ROVODEV_AGENTS_DIR)
}

func (c RovodevCodingAgent) GetProjectSubagentConfigDir(projectDir string) string {
	return filepath.Join(projectDir, ROVODEV_CONFIG_DIR, ROVODEV_AGENTS_DIR)
}

func (c RovodevCodingAgent) GetUserSubagentConfigFilePath(subagentName string) string {
	return filepath.Join(c.GetUserSubagentConfigDir(), subagentName+".md")
}

func (c RovodevCodingAgent) GetProjectSubagentConfigFilePath(projectDir string, subagentName string) string {
	return filepath.Join(c.GetProjectSubagentConfigDir(projectDir), subagentName+".md")
}
