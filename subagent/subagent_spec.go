package subagent

import "github.com/sathish316/pied-piper/config"

type SubagentSpec struct {
	Role                             string
	Description                      string
	Nickname                         string
	TaskLabels                       config.TaskLabelsConfig
	WikiLabels                       config.WikiLabelsConfig
	GeneratedTaskWorkflowDescription string
	GeneratedWikiWorkflowDescription string
	RoleDescription                  string
	Memory                           string
	Model                            string
	RouterModel                      string
}

// Convert SubagentSpec to SubagentConfig - Keeping 2 structs to avoid circular dependencies
func (s *SubagentSpec) ToConfig() *config.SubagentSpecConfig {
	return &config.SubagentSpecConfig{
		Role:                             s.Role,
		Description:                      s.Description,
		Nickname:                         s.Nickname,
		TaskLabels:                       s.TaskLabels,
		WikiLabels:                       s.WikiLabels,
		GeneratedTaskWorkflowDescription: "",
		GeneratedWikiWorkflowDescription: "",
		RoleDescription:                  s.RoleDescription,
		Memory:                           s.Memory,
		Model:                            s.Model,
		RouterModel:                      s.RouterModel,
	}
}
