package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	baseURL    = "https://discord.com/api/webhooks/%s/%s"
	noExtraURL string
)

type Payload struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}

type Hook struct {
	id, token string
}

func (h Hook) SendMessage(payload Payload) error {
	return h.doRequest(http.MethodPost, noExtraURL, payload)
}

func (h Hook) EditMessage(id string, payload Payload) error {
	return h.doRequest(http.MethodPatch, "message/"+id, payload)
}

func (h Hook) doRequest(method, extraURL string, payload Payload) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf(baseURL, h.id, h.token) + extraURL)
	req, err := http.NewRequest(method, fmt.Sprintf(baseURL, h.id, h.token)+extraURL, &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	fmt.Println(err)
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
	return err
}

func NewHook(id, token string) Hook {
	return Hook{
		id:    id,
		token: token,
	}
}
