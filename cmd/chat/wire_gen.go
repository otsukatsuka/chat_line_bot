// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/gomodule/redigo/redis"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/otsukatsuka/chat_line_bot/domain/model"
	"github.com/otsukatsuka/chat_line_bot/domain/repository"
	"github.com/otsukatsuka/chat_line_bot/interface/chat-gpt"
	"github.com/otsukatsuka/chat_line_bot/interface/handler"
	"github.com/otsukatsuka/chat_line_bot/interface/line"
	"github.com/otsukatsuka/chat_line_bot/usecase"
)

// Injectors from wire.go:

// Wire .
func newRouter(ctx context.Context, lineBotClient linebot.Client, chatGPTUrl model.ChatGPTURL, chatGPTApiKey model.ChatGPTApiKey, chatGPTModel model.ChatGPTModel, redisConn redis.Conn) chi.Router {
	lineLine := line.NewLineClient(lineBotClient)
	chatGPTClient := chat_gpt.NewChatGPTClient(chatGPTUrl, chatGPTApiKey, chatGPTModel)
	store := repository.NewStore(redisConn)
	usecaseUsecase := usecase.NewEcho(lineLine, chatGPTClient, store)
	webHookHandler := handler.NewWebHookHandler(lineBotClient, usecaseUsecase)
	router := handler.NewRouter(webHookHandler)
	return router
}
