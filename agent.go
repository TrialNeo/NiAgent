package NiAgent

type Agent struct {
	role   string
	prompt string
	tools  map[string]*NAgentTool
}

// NewAgent 创建一个agent，为其提供工具文档
// role
// prompt
// tools 可用WithTool构建
func NewAgent(role string, prompt string, tools ...*NAgentTool) (a *Agent) {
	a = &Agent{
		role:   role,
		prompt: prompt,
		tools:  make(map[string]*NAgentTool),
	}
	for _, tool := range tools {
		a.tools[tool.name] = tool
	}
	return
}
