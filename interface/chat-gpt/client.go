package chat_gpt

import (
	"bytes"
	"encoding/json"
	"github.com/otsukatsuka/chat_line_bot/domain/model"
	"io"
	"log"
	"net/http"
)

type ChatGPTClient interface {
	Talk(messages model.Messages) (*OpenaiResponse, error)
}

type chatGPTClient struct {
	url    model.ChatGPTURL
	apiKey model.ChatGPTApiKey
	model  model.ChatGPTModel
}

func NewChatGPTClient(
	url model.ChatGPTURL,
	apiKey model.ChatGPTApiKey,
	model model.ChatGPTModel,
) ChatGPTClient {
	return &chatGPTClient{
		url:    url,
		apiKey: apiKey,
		model:  model,
	}
}

type OpenaiRequest struct {
	Model    string         `json:"model"`
	Messages model.Messages `json:"messages"`
}

type OpenaiResponse struct {
	ID      string        `json:"id"`
	Object  string        `json:"object"`
	Created int           `json:"created"`
	Choices model.Choices `json:"choices"`
	Usages  model.Usage   `json:"usage"`
}

func (c *chatGPTClient) Talk(messages model.Messages) (*OpenaiResponse, error) {
	log.Print("Talk")
	requestBody := OpenaiRequest{
		Model:    string(c.model),
		Messages: messages,
	}

	requestJSON, _ := json.Marshal(requestBody)
	log.Print("**********")
	log.Print(requestJSON)
	req, err := http.NewRequest("POST", string(c.url), bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+string(c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("AAAAAAAAAAAA %v", string(body))
	var response OpenaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return &OpenaiResponse{}, nil
	}

	messages = append(messages, model.Message{
		Role:    model.Assistant,
		Content: response.Choices[0].Message.Content,
	})

	return &response, nil
}
