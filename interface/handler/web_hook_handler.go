package handler

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/otsukatsuka/chat_line_bot/usecase"
	"github.com/otsukatsuka/chat_line_bot/usecase/dto"
	"log"
	"net/http"
)

type WebHookHandler interface {
	CallBack(w http.ResponseWriter, req *http.Request)
}

type webHookHandler struct {
	lineBot linebot.Client
	usecase usecase.Usecase
}

func NewWebHookHandler(lineBot linebot.Client, usecase usecase.Usecase) WebHookHandler {
	return &webHookHandler{
		lineBot: lineBot,
		usecase: usecase,
	}
}

func (h *webHookHandler) CallBack(w http.ResponseWriter, req *http.Request) {
	events, err := h.lineBot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			message := dto.Message{
				ReplyToken:  event.ReplyToken,
				LineMessage: event.Message,
			}
			if err = h.usecase.TalkToChatGPT(req.Context(), message); err != nil {
				log.Print(err)
			}
		}
	}
}
