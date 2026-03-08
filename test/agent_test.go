package test

import (
	"NiAgent"
	"testing"
)

func getWeather(location string) string {
	return location + "good weather"
}

func getLocation() string {
	return "GuangDong GuangZhou"
}

func TestFullToolAgent(t *testing.T) {
	NiAgent.NewAgent(
		"weather reporter",
		"帮助用户预测天气",
		NiAgent.WithTool("getWeather", "获取当前天气", getWeather, "location"),
		NiAgent.WithTool("getUserLocation", "获取用户的地址：如广州，上海等", getLocation),
	).ToSchema()
}

func TestLackToolAgent(t *testing.T) {
	NiAgent.NewAgent(
		"weather reporter",
		"帮助用户预测天气",
		NiAgent.WithTool("getWeather", "获取当前天气", getWeather, "location"),
	).ToSchema()
}
