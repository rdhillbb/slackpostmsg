package slackpostmsg

//package main

import (
	"github.com/bluele/slack"
)

func Slkpostmsg(token string, channelName string, messages string) {
	if token == "" {return}
	if channelName == "" {return}
	api := slack.New(token)
	err := api.ChatPostMessage(channelName, messages, nil)
	if err != nil {
		panic(err)
	}
}
func main() {
	Slkpostmsg("xoxb-233290261431-mBepKpXthryR1pWBj9hUdq42", "C6EN7P5U3", "P1 Alert: 174560 Customer:NetMagic http://www.cnn.com")
}
