package NiAgent

import (
	"context"
	"reflect"
)

type NAgentTool struct {
	name string
	ctx  context.Context
	desc string //对于本tool的描述,功能解释
	Tool reflect.Value
	args *toolArgs
}

type toolArgs struct {
	isVariadic bool //最后一个参数是否为可变参数
	num        int
	require    []string
	types      []reflect.Type //每个参数依次的类型
}

// WithTool 创建tool
func WithTool(name, desc string, method any, require ...string) *NAgentTool {
	methodType := reflect.TypeOf(method)
	if methodType.Kind() != reflect.Func {
		//在这乱传函数进来，岂可修
		return nil
	}

	argTypes := make([]reflect.Type, methodType.NumIn())
	if methodType.NumIn() > len(require) {
		return nil
	}
	for i := 0; i < methodType.NumIn(); i++ {
		argTypes[i] = methodType.In(i)
	}
	methodRef := reflect.ValueOf(method)
	return &NAgentTool{
		name: name,
		desc: desc,
		Tool: methodRef,
		args: &toolArgs{
			num:        methodType.NumIn(),
			isVariadic: methodType.IsVariadic(),
			require:    require,
			types:      argTypes,
		},
	}
}
