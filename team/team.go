package team

import (
	"github.com/sathish316/pied-piper/config"
)

type Team struct {
	Name        string
	Description string
	TeamConfig  *config.TeamConfig
}

func (t *Team) GetConfig() *config.TeamConfig {
	return t.TeamConfig
}

func (t *Team) SetConfig(data *config.TeamConfig) {
	t.TeamConfig = data
}
