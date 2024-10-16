package openweather

import (
	"context"
	"net/http"
)

const (
	DefaultUserAgent = "github.com/guythatdrinkscoffee/go-openweather"
)

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

	if c.userAgent == "" {
		c.userAgent = DefaultUserAgent
	}

	return c
}

func WithUserAgent(agent string) ClientOption {
	return func(c *Client) {
		c.userAgent = agent
	}
}

func (c *Client) Weather(ctx context.Context, request Request, dataSet WeatherDataSet) (*WeatherResponse, error) {
	res := &WeatherResponse{}

	var err error

	switch dataSet {
	case CurrentWeatherDataSet:
		err = c.performRequest(ctx, request, string(dataSet), c.token, &res.CurrentWeather)
	case ForecastWeatherDataSet:
		err = c.performRequest(ctx, request, string(dataSet), c.token, &res.ForecastWeather)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) performRequest(ctx context.Context, request requestBuilder, path string, token string, output interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, request.url(path, token), nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", c.userAgent)

	req.Header.Set("Accept-Encoding", "gzip")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	err = validateHttpResponse(res)
	if err != nil {
		return err
	}

	return bind(res, &output)
}
