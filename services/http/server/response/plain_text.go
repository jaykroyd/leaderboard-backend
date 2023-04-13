package response

import app "github.com/byyjoww/leaderboard/services/http/server"

var (
	_ app.Response = (*PlainText)(nil)
)

type PlainText struct {
	content     string
	contentType string
	statusCode  int
	headers     map[string]string
}

func NewPlainText(content string, statusCode int) PlainText {
	return PlainText{
		content:     content,
		contentType: "application/json",
		statusCode:  statusCode,
	}
}

func (h PlainText) GetContent() ([]byte, error) {
	return []byte(h.content), nil
}

func (h PlainText) GetContentType() string {
	return h.contentType
}

func (h PlainText) GetStatusCode() int {
	return h.statusCode
}

func (h PlainText) GetHeaders() map[string]string {
	return h.headers
}
