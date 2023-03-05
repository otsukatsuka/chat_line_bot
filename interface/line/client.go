package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LinetClient interface {
	ReplyMessage(replyToken string, content string) error
}

type lineClient struct {
	client linebot.Client
}

func NewLineClient(client linebot.Client) LinetClient {
	return &lineClient{client: client}
}

func (c *lineClient) ReplyMessage(replyToken string, content string) error {
	response := linebot.NewTextMessage(content)
	if _, err := c.client.ReplyMessage(replyToken, response).Do(); err != nil {
		return err
	}
	return nil
}
