package team

import "path/filepath"

const DEFAULT_CONFIG_DIR = ".pied-piper"
const DEFAULT_CONFIG_FILE = "config.yml"

type Team struct {
	Name string
	Description string
	TeamConfig *TeamConfig
}

func (t *Team) GetConfigFilePath() string {
	return filepath.Join(t.GetConfigDir(), t.configFile())
}

func (t *Team) GetConfigDir() string {
	return t.TeamConfig.Path
}

func (t *Team) configFile() string {
	return t.TeamConfig.File
}

func (t *Team) GetConfigData() map[string]interface{} {
	return t.TeamConfig.Data
}

func (t *Team) SetConfigData(data map[string]interface{}) {
	t.TeamConfig.Data = data
}