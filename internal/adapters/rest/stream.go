package rest

import (
	"encoding/json"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type StreamRest struct {
	body []byte
}

// Parse implements ports.StreamRest.
func (s *StreamRest) Parse(body any) error {
	return json.Unmarshal(s.body, body)
}

func NewStreamRest(body []byte) ports.StreamRest {

	return &StreamRest{body: body}
}
