package bilibili

import (
	"net/http"
)

type Client struct {
	AccessKey       string
	AccessKeySecret string
	Client          *http.Client
}

func NewClient(key, secret string) *Client {
	return &Client{
		AccessKey:       key,
		AccessKeySecret: secret,
		Client:          http.DefaultClient,
	}
}
