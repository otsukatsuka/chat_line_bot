package main

import (
	"context"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/otsukatsuka/chat_line_bot/config"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	lineConfig, err := config.NewLineMessageConfig()
	if err != nil {
		log.Print(err)
	}
	lineClient, err := linebot.New(lineConfig.SECRET, lineConfig.TOKEN)
	if err != nil {
		log.Print(err)
	}
	chatGPTConfig, err := config.NewChatGPTConfig()
	if err != nil {
		log.Print(err)
	}
	apiConfig, err := config.NewAPIConfig()
	if err != nil {
		log.Print(err)
	}
	r := newRouter(
		ctx,
		*lineClient,
		chatGPTConfig.URL,
		chatGPTConfig.APIKEY,
		chatGPTConfig.MODEL,
	)
	addr := fmt.Sprintf(":%v", apiConfig.PORT)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("failed run server")
	}
}
