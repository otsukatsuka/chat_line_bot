package model

type ChatGPTURL string
type ChatGPTApiKey string
type ChatGPTModel string

type ChatGPTRole string

const (
	System    ChatGPTRole = "system"
	Assistant             = "assistant"
	User                  = "user"
)

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Choices []Choice

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Messages []Message

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
