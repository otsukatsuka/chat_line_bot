//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/otsukatsuka/chat_line_bot/domain/model"
	"github.com/otsukatsuka/chat_line_bot/domain/repository"
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
	redisConn redis.Conn,
) chi.Router {
	wire.Build(
		line.NewLineClient,
		chat_gpt.NewChatGPTClient,
		repository.NewStore,
		usecase.NewEcho,
		handler.NewWebHookHandler,
		handler.NewRouter,
	)
	return &chi.Mux{}
}
