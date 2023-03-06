package usecase

import (
	"context"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/otsukatsuka/chat_line_bot/domain/model"
	"github.com/otsukatsuka/chat_line_bot/domain/repository"
	chat_gpt "github.com/otsukatsuka/chat_line_bot/interface/chat-gpt"
	"github.com/otsukatsuka/chat_line_bot/interface/line"
	"github.com/otsukatsuka/chat_line_bot/usecase/dto"
	"log"
)

type Usecase interface {
	TalkToChatGPT(ctx context.Context, message dto.Message) error
}

type usecase struct {
	lineClient      line.Line
	chatGPTClient   chat_gpt.ChatGPTClient
	storeRepository repository.Store
}

func (u usecase) TalkToChatGPT(ctx context.Context, message dto.Message) error {
	switch msg := message.LineMessage.(type) {
	case *linebot.TextMessage:
		currentMessage := model.Message{
			Role:    model.User,
			Content: msg.Text,
		}
		pastMessages, err := u.storeRepository.GetMessages()
		log.Print(pastMessages)
		if err != nil && err != redis.ErrNil {
			log.Print(err)
			return err
		}
		var chatGPTMessages model.Messages
		if err == redis.ErrNil {
			chatGPTMessages = model.Messages{currentMessage}
		} else {
			chatGPTMessages = append(pastMessages, currentMessage)
		}
		res, messages, err := u.chatGPTClient.Talk(chatGPTMessages)
		if err != nil {
			log.Print(err)
			return err
		}
		b, err := json.Marshal(messages)
		if err := u.storeRepository.SetMessages(b); err != nil {
			log.Print(err)
			return err
		}
		if err != nil {
			log.Print(err)
			return err
		}
		if err := u.lineClient.ReplyMessage(message.ReplyToken, res.Choices[0].Message.Content); err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

func NewEcho(
	lineClient line.Line,
	chatCPTClient chat_gpt.ChatGPTClient,
	storeRepository repository.Store,
) Usecase {
	return &usecase{
		lineClient:      lineClient,
		chatGPTClient:   chatCPTClient,
		storeRepository: storeRepository,
	}
}
