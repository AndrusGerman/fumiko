package ollama

import "github.com/AndrusGerman/fumiko/internal/core/ports"

type ollama struct {
}

func New() ports.LLM {
	var llm = new(ollama)
	return llm
}
