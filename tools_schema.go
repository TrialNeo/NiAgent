package NiAgent

// ToToolSchema 把工具转换为schema
func (t *NAgentTool) ToToolSchema() (schema map[string]any) {
	parameters := make(map[string]any)
	for i := range t.args.require {
		parameters[t.args.require[i]] = map[string]string{
			"type": t.args.types[i].String(),
		}
	}
	schema = map[string]any{
		"role": t.name,
		"desc": t.desc,
		"parameters": map[string]any{
			"type":       "object",
			"require":    t.args.require,
			"properties": parameters,
		},
	}
	return
}
