package NiAgent

import (
	"encoding/json"
	"fmt"
)

type Agent struct {
	role   string
	prompt string
	tools  []*NAgentTool
}

// NewAgent 创建一个agent，为其提供工具文档
// role
// prompt
// tools 可用WithTool构建
func NewAgent(role string, prompt string, tools ...*NAgentTool) *Agent {

	return &Agent{
		role:   role,
		prompt: prompt,
		tools:  tools,
	}

}

func (a *Agent) ToSchema() {
	toolsSchema := make([]map[string]any, len(a.tools))
	for i := 0; i < len(a.tools); i++ {
		toolsSchema[i] = a.tools[i].ToToolSchema()
	}
	schema := map[string]any{
		"role":          a.role,
		"system_prompt": prompt,
		"user_prompt":   a.prompt,
		"tools":         toolsSchema,
		"rules":         rules,
	}
	marshal, _ := json.Marshal(schema)
	fmt.Println(string(marshal))
}
