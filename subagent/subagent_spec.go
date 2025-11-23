package subagent

import "github.com/sathish316/pied-piper/config"

type SubagentSpec struct {
	Role		string
	Description string
	Nickname	string
	TaskLabels  config.TaskLabelsConfig
	WikiLabels  config.WikiLabelsConfig
	WorkflowDescription string
	RoleDescription string
	Memory string
}

// Convert SubagentSpec to SubagentConfig - Keeping 2 structs to avoid circular dependencies
func (s *SubagentSpec) ToConfig() *config.SubagentSpecConfig {
	return &config.SubagentSpecConfig{
		Role: s.Role,
		Description: s.Description,
		Nickname: s.Nickname,
		TaskLabels: s.TaskLabels,
		WikiLabels: s.WikiLabels,
		WorkflowDescription: s.WorkflowDescription,
		RoleDescription: s.RoleDescription,
		Memory: s.Memory,
	}
}