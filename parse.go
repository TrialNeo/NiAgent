package NiAgent

import (
	"encoding/json"
)

func (a *Agent) Parse(data string) {
	//	假设都是正常数据
	action := map[string]any{}
	if err := json.Unmarshal([]byte(data), &action); err != nil {
		return
	}
	method := action["action"].(string)
	if method == "" {
		return
	}
	actionInput := action["action_input"].(map[string]any)
	params := make([]any, len(actionInput))
	count := 0
	for _, param := range actionInput {
		params[count] = param
		count++
	}
	a.tools[method].Call(params...)
}
