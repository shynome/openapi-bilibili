package bilibili

import (
	"net/http"
)

type Client struct {
	Endpoint        string
	AccessKey       string
	AccessKeySecret string
	Client          *http.Client
}

func NewClient(key, secret string) *Client {
	return &Client{
		Endpoint:        "https://live-open.biliapi.com",
		AccessKey:       key,
		AccessKeySecret: secret,
		Client:          http.DefaultClient,
	}
}
