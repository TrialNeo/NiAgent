package test

import (
	"NiAgent"
	"testing"
)

func getWeather(location string) string {
	return "good weather"
}

func TestNoneAgent(t *testing.T) {
	NiAgent.NewAgent(
		"weather reporter",
		"你是一位天气预报员",
		NiAgent.WithTool("getWeather", "获取当前天气", getWeather, "location"),
	)
}
