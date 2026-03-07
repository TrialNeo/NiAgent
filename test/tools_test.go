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

func F2(a ...interface{}) {
	fmt.Println(a...)
}

func TestTo(t *testing.T) {
	NiAgent.WithTool("F2", "这是一个描述", F2).Call([]any{1, 2, "3"})

}
