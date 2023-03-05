package dto

import "github.com/line/line-bot-sdk-go/v7/linebot"

type Message struct {
	ReplyToken  string
	LineMessage linebot.Message
}
