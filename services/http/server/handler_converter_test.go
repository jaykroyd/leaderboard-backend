package server_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	lib "github.com/byyjoww/leaderboard/services/http"
	app "github.com/byyjoww/leaderboard/services/http/server"
	"github.com/byyjoww/leaderboard/services/http/server/response"
	"github.com/byyjoww/leaderboard/services/http/test"
	"github.com/stretchr/testify/assert"
)

type TestHandler struct {
	Method   string
	Path     string
	Response app.Response
}

func (h *TestHandler) GetMethod() string {
	return h.Method
}

func (h *TestHandler) GetPath() string {
	return h.Path
}

func (h *TestHandler) Handle(logger lib.Logger, r *http.Request) app.Response {
	return h.Response
}

func TestHandlerConverter(t *testing.T) {
	var (
		method   = http.MethodGet
		path     = "/test"
		text     = "OK"
		response = response.NewJsonStatusOK(text)
	)

	handler := app.NewHandlerConverter(
		&TestHandler{
			Method:   method,
			Path:     path,
			Response: response,
		},
	)

	assert.Equal(t, handler.GetMethod(), method)
	assert.Equal(t, handler.GetPath(), path)

	req := httptest.NewRequest(handler.GetMethod(), handler.GetPath(), nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(test.NewNullLogger(), w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	expected, _ := json.Marshal(text)
	assert.Nil(t, err)
	assert.Equal(t, data, expected)
}
