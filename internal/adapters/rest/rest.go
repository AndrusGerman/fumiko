package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type rest struct {
}

// Post implements ports.Rest.
func (r *rest) Post(url string, body any, out any) error {
	var err error
	var bodyByte []byte
	var resp *http.Response
	if bodyByte, err = json.Marshal(body); err != nil {
		return err
	}

	if resp, err = http.Post(url, "application/json", bytes.NewReader(bodyByte)); err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(out)
}

func New() ports.Rest {
	return &rest{}
}
