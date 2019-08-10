package auspost

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	DefaultApiBaseUrl = "https://digitalapi.auspost.com.au"
)

func NewDefaultClient(apiKey string) *Client {
	return NewClient(apiKey, DefaultApiBaseUrl)
}

func NewClient(apiKey string, baseUrl string) *Client {
	client := resty.New()
	client.SetHeader("auth-key", apiKey)

	return &Client{
		restClient: client,
		baseUrl:    baseUrl,
	}
}

type Client struct {
	restClient *resty.Client
	baseUrl    string
	trace      bool
}

func (c *Client) getUrl(uri string) string {
	return fmt.Sprintf("%s/%s", c.baseUrl, uri)
}

func (c *Client) EnableTrace() *Client {
	c.trace = true

	return c
}
