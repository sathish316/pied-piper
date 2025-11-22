package subagent

type SubAgentSpec struct {
	Name string
	Description string
	Prompt string
	Tools []string
	WorkflowDescription string
	BehaviourDescription string
}
