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

func spewTrace(resp *resty.Response, err error) {
	fmt.Println("AusPost API Request")
	fmt.Println("===================")
	fmt.Println()
	fmt.Println("Request Info:")
	fmt.Println("URL      :", resp.Request.URL)
	fmt.Println()
	fmt.Println("Response Info:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status     :", resp.Status())
	fmt.Println("Time       :", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Body       :\n", resp)
	fmt.Println()

	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("DNSLookup    :", ti.DNSLookup)
	fmt.Println("ConnTime     :", ti.ConnTime)
	fmt.Println("TLSHandshake :", ti.TLSHandshake)
	fmt.Println("ServerTime   :", ti.ServerTime)
	fmt.Println("ResponseTime :", ti.ResponseTime)
	fmt.Println("TotalTime    :", ti.TotalTime)
	fmt.Println("IsConnReused :", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime :", ti.ConnIdleTime)
	fmt.Println()
}
