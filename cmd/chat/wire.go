//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/otsukatsuka/chat_line_bot/domain/model"
	chat_gpt "github.com/otsukatsuka/chat_line_bot/interface/chat-gpt"
	"github.com/otsukatsuka/chat_line_bot/interface/handler"
	"github.com/otsukatsuka/chat_line_bot/interface/line"
	"github.com/otsukatsuka/chat_line_bot/usecase"
)

// Wire .
func newRouter(
	ctx context.Context,
	lineBotClient linebot.Client,
	chatGPTUrl model.ChatGPTURL,
	chatGPTApiKey model.ChatGPTApiKey,
	chatGPTModel model.ChatGPTModel,
) chi.Router {
	wire.Build(
		line.NewLineClient,
		chat_gpt.NewChatGPTClient,
		usecase.NewEcho,
		handler.NewWebHookHandler,
		handler.NewRouter,
	)
	return &chi.Mux{}
}
