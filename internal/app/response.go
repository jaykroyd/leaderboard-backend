package app

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	GetContent() ([]byte, error)
	GetContentType() string
	GetStatusCode() int
}

type JsonResponse struct {
	Content     interface{}
	ContentType string
	StatusCode  int
}

func NewJsonResponse(content interface{}, statusCode int) JsonResponse {
	return JsonResponse{
		Content:     content,
		ContentType: "application/json",
		StatusCode:  statusCode,
	}
}

func NewForbidden() JsonResponse {
	return NewJsonResponse(map[string]string{
		"error": "forbidden",
	}, http.StatusForbidden)
}

func NewBadRequest(err error) JsonResponse {
	return NewJsonResponse(map[string]string{
		"error": err.Error(),
	}, http.StatusBadRequest)
}

func NewInternalServerError(err error) JsonResponse {
	return NewJsonResponse(map[string]string{
		"error": err.Error(),
	}, http.StatusInternalServerError)
}

func NewStatusOK(content interface{}) JsonResponse {
	return NewJsonResponse(content, http.StatusOK)
}

func (h JsonResponse) GetContent() ([]byte, error) {
	return json.Marshal(h.Content)
}

func (h JsonResponse) GetContentType() string {
	return h.ContentType
}

func (h JsonResponse) GetStatusCode() int {
	return h.StatusCode
}

type PlainTextResponse struct {
	Content     string
	ContentType string
	StatusCode  int
}

func NewPlainTextResponse(content string, statusCode int) PlainTextResponse {
	return PlainTextResponse{
		Content:     content,
		ContentType: "application/json",
		StatusCode:  statusCode,
	}
}

func (h PlainTextResponse) GetContent() ([]byte, error) {
	return []byte(h.Content), nil
}

func (h PlainTextResponse) GetContentType() string {
	return h.ContentType
}

func (h PlainTextResponse) GetStatusCode() int {
	return h.StatusCode
}
