package NiAgent

type Agent struct {
	name   string
	prompt string
	tools  []*NAgentTool
}

// NewAgent 创建一个agent，为其提供工具文档
func NewAgent(role string, prompt string, tools ...*NAgentTool) *Agent {
	return &Agent{
		name:   role,
		prompt: prompt,
		tools:  tools,
	}
}
