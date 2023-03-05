package usecase

import (
	"context"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/otsukatsuka/chat_line_bot/domain/model"
	chat_gpt "github.com/otsukatsuka/chat_line_bot/interface/chat-gpt"
	"github.com/otsukatsuka/chat_line_bot/interface/line"
	"github.com/otsukatsuka/chat_line_bot/usecase/dto"
	"log"
)

type Usecase interface {
	TalkToChatGPT(ctx context.Context, message dto.Message) error
}

type usecase struct {
	lineClient    line.LinetClient
	chatGPTClient chat_gpt.ChatGPTClient
}

func (u usecase) TalkToChatGPT(ctx context.Context, message dto.Message) error {
	log.Print("TalkToChatGPT")
	switch msg := message.LineMessage.(type) {
	case *linebot.TextMessage:
		chatGPTMessages := model.Messages{
			model.Message{
				Role:    model.User,
				Content: msg.Text,
			},
		}
		log.Print(chatGPTMessages)
		log.Print("u.chatGPTClient.Talk")
		res, err := u.chatGPTClient.Talk(chatGPTMessages)
		if err != nil {
			log.Print(err)
			return err
		}
		log.Print("ReplyMessage")
		if err := u.lineClient.ReplyMessage(message.ReplyToken, res.Choices[0].Message.Content); err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

func NewEcho(lineClient line.LinetClient, chatCPTClient chat_gpt.ChatGPTClient) Usecase {
	return &usecase{
		lineClient:    lineClient,
		chatGPTClient: chatCPTClient,
	}
}
