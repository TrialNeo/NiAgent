package NiAgent

import (
	"context"
	"reflect"
)

type NAgentTool struct {
	ctx  context.Context
	desc string //对于本tool的描述,功能解释
	Tool reflect.Value
	args *toolArgs
}

type toolArgs struct {
	isVariadic bool //最后一个参数是否为可变参数
	num        int
	types      []reflect.Type //每个参数依次的类型
}

// WithTool 创建tool
func WithTool(name, desc string, method any) *NAgentTool {
	methodType := reflect.TypeOf(method)
	if methodType.Kind() != reflect.Func {
		//在这乱传函数进来，岂可修
		return nil
	}

	argTypes := make([]reflect.Type, methodType.NumIn())
	for i := 0; i < methodType.NumIn(); i++ {
		argTypes[i] = methodType.In(i)
	}

	return &NAgentTool{
		desc: desc,
		Tool: reflect.ValueOf(method),
		args: &toolArgs{
			num:        methodType.NumIn(),
			isVariadic: methodType.IsVariadic(),
			types:      argTypes,
		},
	}
}

// Call agent调用工具，但是只能传递string,int,uint...,bool,float,float64之类的基础类型
func (tool *NAgentTool) Call(args []any) (err error) {
	if tool.args.isVariadic {

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
