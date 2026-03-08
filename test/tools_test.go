package test

import (
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

}
