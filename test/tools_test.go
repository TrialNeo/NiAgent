package test

import (
	"NiAgent"
	"fmt"
	"testing"
	"time"
)

func Time() {
	fmt.Println(time.Now())
}

func TestTo(t *testing.T) {
	NiAgent.WithTool("获取时间", "这是一个描述", Time).Call(nil)
}
