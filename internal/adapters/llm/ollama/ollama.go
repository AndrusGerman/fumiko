package ollama

import (
	"strings"
	"sync"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type ollama struct {
	rest   ports.Rest
	model  string
	memory []*message
	mt     *sync.Mutex
}

func New(rest ports.Rest) ports.LLM {
	var llm = new(ollama)
	llm.rest = rest
	llm.mt = new(sync.Mutex)
	llm.model = "gemma2:latest"
	return llm
}

func (o *ollama) BasicQuest(text string) string {
	var m = newMessage("user", strings.TrimSpace(text))
	o.memory = append(o.memory, m)

	var request = newOllamaRequest(o.model, o.memory)
	var err error
	var response = new(messageResponse)

	if err = o.newRequest(request, response); err != nil {
		return err.Error()
	}

	o.memory = append(o.memory, newMessage(response.Message.Role, strings.TrimSpace(response.Message.Content)))
	return response.Message.Content
}

func (o *ollama) newRequest(request *ollamaRequest, response *messageResponse) error {
	o.mt.Lock()
	defer o.mt.Unlock()
	return o.rest.Post("http://localhost:11434/api/chat", request, &response)
}
