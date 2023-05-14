package model

import (
	"time"
)

type CreateRequest struct {
	URL         string        `json:"url"`
	ShorternURL string        `json:"shortern_url"`
	Expiry      time.Duration `json:"expiry"`
}

type Response struct {
	URL         string        `json:"url" valid:"url"`
	ShorternURL string        `json:"shortern_url"`
	Expiry      time.Duration `json:"expiry"`
	UserId      string        `json:"userid,omitempty"`
}

type Status struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

var M map[string]Response
