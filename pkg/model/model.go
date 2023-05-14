package model

import (
	"errors"
	"log"
	"time"

	"github.com/padhitheboss/url-shortner-go/pkg/config"
	"github.com/redis/go-redis/v9"
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
var DB *redis.Client

func init() {
	DB = config.CreateClient(0)
}

func StoreURL(body *CreateRequest) error {
	val, _ := DB.Get(config.Ctx, body.ShorternURL).Result()
	if val != "" {
		return errors.New("the shorturl is already in  use")
	}
	err := DB.Set(config.Ctx, body.ShorternURL, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}

func GetURL(shurl string) (string, error) {
	url, err := DB.Get(config.Ctx, shurl).Result()
	if url == "" {
		return url, errors.New("the shorturl has expired")
	}
	return url, err
}
