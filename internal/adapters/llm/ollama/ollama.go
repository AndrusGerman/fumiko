package ollama

import (
	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type ollama struct {
	rest  ports.Rest
	model string
}

func New(rest ports.Rest) ports.LLM {
	var llm = new(ollama)
	llm.rest = rest
	llm.model = "gemma2:latest"
	return llm
}

func (o *ollama) BasicQuest(text string) string {
	var m = newMessage("user", text)
	var request = newOllamaRequest(o.model, []*message{m})
	var err error
	var response = new(messageResponse)

	if err = o.rest.Post("http://localhost:11434/api/chat", request, &response); err != nil {
		return err.Error()
	}

	return response.Message.Content
}
