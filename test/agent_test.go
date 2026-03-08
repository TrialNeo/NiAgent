package test

import (
	"NiAgent"
	"fmt"
	"testing"
)

func getWeather(location string) string {
	fmt.Printf("[Agent Action]:getWeather(%s)\n", location)
	return location + "good weather"
}

func getLocation() string {
	fmt.Printf("[Agent Action]:getLocation()\n")
	return "GuangDong GuangZhou"
}

func TestFullToolAgent(t *testing.T) {
	a := NiAgent.NewAgent(
		"weather reporter",
		"帮助用户预测天气",
		NiAgent.WithTool("getWeather", "获取当前天气", getWeather, "location"),
		NiAgent.WithTool("getUserLocation", "获取用户的地址：如广州，上海等", getLocation),
	)
	a.Parse(`{"thought": "用户提供了位置信息（广州天河区），现在可以查询该位置的天气", "action": "getWeather", "action_input": {"location": "广州天河区"}}`)
}

func TestLackToolAgent(t *testing.T) {
	NiAgent.NewAgent(
		"weather reporter",
		"帮助用户预测天气",
		NiAgent.WithTool("getWeather", "获取当前天气", getWeather, "location"),
	).ToSchema()
}
