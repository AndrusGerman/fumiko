package ollama

import (
	"log"
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

	var request = newOllamaRequest(o.model, memory, false)
	var err error
	var response *messageResponse

	if response, err = o.newRequest(request); err != nil {
		return "", nil
	}

	return response.Message.Content, nil
}

func (o *ollama) newMessages(base []*domain.Message) []*message {
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
	return messages
}

func (o *ollama) Quest(base []*domain.Message, text string) (*domain.Message, error) {
	var messages = o.newMessages(base)
	var response *messageResponse

	var m = newMessage("user", strings.TrimSpace(text))
	messages = append(messages, m)

	var request = newOllamaRequest(o.model, messages, false)
	var err error

	if response, err = o.newRequest(request); err != nil {
		return nil, err
	}

	return domain.NewMessage(strings.TrimSpace(response.Message.Content), domain.AssistantRoleID), nil
}

func (o *ollama) QuestParts(base []*domain.Message, text string, partsSize int) (<-chan *domain.Message, error) {
	var messages = o.newMessages(base)
	//var response *messageResponse
	var messageResponseStream <-chan *messageResponse
	var messageStream = make(chan *domain.Message)
	var err error

	var m = newMessage("user", strings.TrimSpace(text))
	messages = append(messages, m)

	var request = newOllamaRequest(o.model, messages, true)

	if messageResponseStream, err = o.newRequestStream(request); err != nil {
		return nil, err
	}

	go func() {
		var text string
		for messageResponse := range messageResponseStream {
			text += messageResponse.Message.Content

			if len(text) > partsSize || messageResponse.Done {
				messageStream <- domain.NewMessage(strings.TrimSpace(text), domain.AssistantRoleID)
				text = ""
			}
		}

		close(messageStream)
	}()

	return messageStream, err
}

func (o *ollama) newRequest(request *ollamaRequest) (*messageResponse, error) {
	o.mt.Lock()
	defer o.mt.Unlock()

	var data, err = o.rest.Stream("http://localhost:11434/api/chat", request)
	if err != nil {
		return nil, err
	}

	var response = new(messageResponse)
	dataRaw := <-data

	if err = dataRaw.Parse(response); err != nil {
		return nil, err
	}

	return response, nil
}

func (o *ollama) newRequestStream(request *ollamaRequest) (<-chan *messageResponse, error) {
	o.mt.Lock()
	defer o.mt.Unlock()

	var streamData, err = o.rest.Stream("http://localhost:11434/api/chat", request)
	if err != nil {
		return nil, err
	}

	var chanMessageResponse = make(chan *messageResponse)

	go func() {
		for data := range streamData {
			var response = new(messageResponse)

			if err = data.Parse(response); err != nil {
				log.Println("err: ", err)
				continue
			}

			chanMessageResponse <- response
		}

		close(chanMessageResponse)
	}()

	return chanMessageResponse, nil
}
