package team

import "github.com/sathish316/pied-piper/team"

type SubAgentExporter interface {
	Export(source string, team *team.Team, subagentName string) error
}