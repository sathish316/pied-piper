package subagent

type SubAgentConfig struct {
	Name string
	Description string
	Prompt string
	Tools []string
	TaskLabels map[string][]string
	WikiLabels map[string][]string
}