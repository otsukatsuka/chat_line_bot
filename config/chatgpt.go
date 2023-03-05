package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/otsukatsuka/chat_line_bot/domain/model"
)

type ChatGPTConfig struct {
	APIKEY model.ChatGPTApiKey `envconfig:"CHAT_GPT_API_KEY" default:""`
	URL    model.ChatGPTURL    `envconfig:"CHAT_GPT_URL" default:""`
	MODEL  model.ChatGPTModel  `envconfig:"CHAT_GPT_MODEL" default:""`
}

func NewChatGPTConfig() (ChatGPTConfig, error) {
	var chatGPTConfig ChatGPTConfig
	err := envconfig.Process("", &chatGPTConfig)
	return chatGPTConfig, err
}
