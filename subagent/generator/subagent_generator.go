package subagent

import "github.com/sathish316/pied-piper/team"
import "github.com/sathish316/pied-piper/subagent"

type SubAgentGenerator interface {
	Generate(team *team.Team, role string, nickname string) (subagent.SubAgentSpec, error)
}