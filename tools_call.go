package NiAgent

import (
	"fmt"
	"reflect"
)

// Call agent调用工具，但是只能传递string,int,uint...,bool,float,float64之类的基础类型
func (tool *NAgentTool) Call(args ...any) (err error) {
	if tool.args.isVariadic {
		minArgs := tool.args.num - 1
		if len(args) < minArgs {
			return fmt.Errorf("too few arguments for variadic function")
		}
		parameters := make([]reflect.Value, tool.args.num)
		for i := 0; i < minArgs; i++ {
			v := reflect.ValueOf(args[i])
			if !v.Type().AssignableTo(tool.args.types[i]) {
				return fmt.Errorf("arg %d: %v not assignable to %v", i, v.Type(), tool.args.types[i])
			}
			parameters[i] = v
		}
		sliceType := tool.args.types[minArgs]
		elemType := sliceType.Elem()
		vSlice := reflect.MakeSlice(sliceType, 0, len(args)-minArgs)
		for i := minArgs; i < len(args); i++ {
			v := reflect.ValueOf(args[i])
			if !v.Type().AssignableTo(elemType) {
				return fmt.Errorf("variadic arg %d: %v not assignable to %v", i-minArgs, v.Type(), elemType)
			}
			vSlice = reflect.Append(vSlice, v)
		}
		parameters[minArgs] = vSlice

		tool.Tool.CallSlice(parameters)
	} else {
		if len(args) != tool.args.num {
			return
		}
		parameters := make([]reflect.Value, tool.args.num)
		for i := 0; i < tool.args.num; i++ {
			parameters[i] = reflect.ValueOf(args[i])
			if parameters[i].Type() != tool.args.types[i] {
				return
			}
		}
		tool.Tool.Call(parameters)
	}
	return
}
