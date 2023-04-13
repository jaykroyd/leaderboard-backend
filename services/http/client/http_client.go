package client

import (
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RetryHTTPClient struct {
	client HTTPClient
	config RetryConfig
}

type RetryConfig struct {
	MaxRetries     int
	Delay          time.Duration
	RetryCondition func(statusCode int, err error) bool
}

func NewRetryHTTPClient(client HTTPClient, config RetryConfig) *RetryHTTPClient {
	return &RetryHTTPClient{
		client: client,
		config: config,
	}
}

func (c *RetryHTTPClient) Do(req *http.Request) (*http.Response, error) {
	for r := 0; ; r++ {
		response, err := c.client.Do(req)
		if !c.config.RetryCondition(response.StatusCode, err) || r >= c.config.MaxRetries {
			return response, err
		}

		select {
		case <-time.After(c.config.Delay):
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
	}
}
