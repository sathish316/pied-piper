package team

import (
	"github.com/sathish316/pied-piper/config"
	"github.com/sathish316/pied-piper/team"
)

type SubAgentGenerator interface {
	Generate(team *team.Team, role string, nickname string) (config.SubagentConfig, error)
}