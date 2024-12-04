package ollama

import (
	"strings"
	"sync"

	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type ollama struct {
	rest  ports.Rest
	model string
	mt    *sync.Mutex
}

func New(rest ports.Rest) ports.LLM {
	var llm = new(ollama)
	llm.rest = rest
	llm.mt = new(sync.Mutex)
	llm.model = "gemma2:latest"
	return llm
}

func (o *ollama) BasicQuest(text string) (string, error) {
	var memory []*message
	var m = newMessage("user", strings.TrimSpace(text))

	memory = append(memory, m)

	var request = newOllamaRequest(o.model, memory)
	var err error
	var response = new(messageResponse)

	if err = o.newRequest(request, response); err != nil {
		return "", nil
	}

	return response.Message.Content, nil
}

func (o *ollama) Quest(base []*domain.Message, text string) (*domain.Message, error) {

	var messages = make([]*message, len(base))

	for i := range base {
		if base[i].RoleID == domain.AssistantRoleID {
			messages[i] = newMessage("assistant", base[i].Content)
		}

		if base[i].RoleID == domain.UserRoleID {
			messages[i] = newMessage("user", base[i].Content)
		}

		if base[i].RoleID == domain.SystemRoleID {
			messages[i] = newMessage("system", base[i].Content)
		}
	}

	var m = newMessage("user", strings.TrimSpace(text))
	messages = append(messages, m)

	var request = newOllamaRequest(o.model, messages)
	var err error
	var response = new(messageResponse)

	if err = o.newRequest(request, response); err != nil {
		return nil, err
	}

	return domain.NewMessage(strings.TrimSpace(response.Message.Content), domain.AssistantRoleID), nil
}

func (o *ollama) newRequest(request *ollamaRequest, response *messageResponse) error {
	o.mt.Lock()
	defer o.mt.Unlock()
	return o.rest.Post("http://localhost:11434/api/chat", request, &response)
}
