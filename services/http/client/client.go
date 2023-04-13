package client

import (
	"fmt"
	"net/http"
)

type Client interface {
	Get(endpoint string) *HttpRequest
	Post(endpoint string) *HttpRequest
	Put(endpoint string) *HttpRequest
	Patch(endpoint string) *HttpRequest
	Delete(endpoint string) *HttpRequest
	NewRequest(method string, endpoint string) *HttpRequest
}

type ClientImpl struct {
	client  HTTPClient
	decoder Decoder
	baseUrl string
	auth    Auth
}

func New(client HTTPClient, decoder Decoder) *ClientImpl {
	return &ClientImpl{
		client:  client,
		decoder: decoder,
	}
}

func (c *ClientImpl) WithBaseUrl(baseUrl string) *ClientImpl {
	c.baseUrl = baseUrl
	return c
}

func (c *ClientImpl) WithAuth(auth Auth) *ClientImpl {
	c.auth = auth
	return c
}

func (c *ClientImpl) Get(endpoint string) *HttpRequest {
	return c.NewRequest(http.MethodGet, endpoint)
}

func (c *ClientImpl) Post(endpoint string) *HttpRequest {
	return c.NewRequest(http.MethodPost, endpoint)
}

func (c *ClientImpl) Put(endpoint string) *HttpRequest {
	return c.NewRequest(http.MethodPut, endpoint)
}

func (c *ClientImpl) Patch(endpoint string) *HttpRequest {
	return c.NewRequest(http.MethodPatch, endpoint)
}

func (c *ClientImpl) Delete(endpoint string) *HttpRequest {
	return c.NewRequest(http.MethodDelete, endpoint)
}

func (c *ClientImpl) NewRequest(method string, endpoint string) *HttpRequest {
	req := &HttpRequest{
		client:      c.client,
		decoder:     c.decoder,
		ctx:         nil,
		method:      method,
		url:         fmt.Sprintf(c.baseUrl + endpoint),
		contentType: "",
		auth:        c.auth,
		body:        nil,
		query:       map[string]string{},
		headers:     map[string]string{},
	}
	return req
}
