package team

import "github.com/sathish316/pied-piper/team"

type SubAgentImporter interface {
	Import(source string, team *team.Team, subagentName string) error
}