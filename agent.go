package NiAgent

type Agent struct {
	prompt string
	tools  []*NAgentTool
}

// NewAgent 创建一个agent，为其提供工具文档
//func NewAgent() *Agent {
//
//}
