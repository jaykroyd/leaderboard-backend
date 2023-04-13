package response

import (
	"encoding/json"
	"net/http"

	app "github.com/byyjoww/leaderboard/services/http/server"
)

var (
	_ app.Response = (*Json)(nil)
)

type Json struct {
	content     interface{}
	contentType string
	statusCode  int
	headers     map[string]string
}

func NewJson(statusCode int, headers map[string]string, content interface{}) Json {
	return Json{
		content:     content,
		contentType: "application/json",
		statusCode:  statusCode,
		headers:     map[string]string{},
	}
}

func (h Json) GetContent() ([]byte, error) {
	return json.Marshal(h.content)
}

func (h Json) GetContentType() string {
	return h.contentType
}

func (h Json) GetStatusCode() int {
	return h.statusCode
}

func (h Json) GetHeaders() map[string]string {
	return h.headers
}

// ------------- Status Codes -------------

func NewJsonForbidden() Json {
	return NewJson(http.StatusForbidden, map[string]string{}, map[string]string{
		"error": "forbidden",
	})
}

func NewJsonUnauthorized() Json {
	return NewJson(http.StatusUnauthorized, map[string]string{
		"WWW-Authenticate": `Basic realm="restricted", charset="UTF-8"`,
	}, map[string]string{
		"error": "unauthorized",
	})
}

func NewJsonBadRequest(err error) Json {
	return NewJson(http.StatusBadRequest, map[string]string{}, map[string]string{
		"error": err.Error(),
	})
}

func NewJsonInternalServerError(err error) Json {
	return NewJson(http.StatusInternalServerError, map[string]string{}, map[string]string{
		"error": err.Error(),
	})
}

func NewJsonStatusOK(content interface{}) Json {
	return NewJson(http.StatusOK, map[string]string{}, content)
}
