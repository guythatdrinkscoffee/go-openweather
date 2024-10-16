package openweather

import "net/http"

type Client struct {
	token      string
	userAgent  string
	httpClient *http.Client
}

type ClientOption func(c *Client)

func NewClient(apiToken string, opts ...ClientOption) *Client {
	c := &Client{token: apiToken, httpClient: http.DefaultClient}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithUserAgent(agent string) ClientOption {
	return func(c *Client) {
		c.userAgent = agent
	}
}
