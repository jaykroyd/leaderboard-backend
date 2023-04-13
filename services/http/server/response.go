package server

import (
	"fmt"
	"net/http"

	lib "github.com/byyjoww/leaderboard/services/http"
)

type Response interface {
	GetContent() ([]byte, error)
	GetContentType() string
	GetStatusCode() int
	GetHeaders() map[string]string
}

func Respond(logger lib.Logger, w http.ResponseWriter, resp Response) {
	var (
		contentType = resp.GetContentType()
		headers     = resp.GetHeaders()
		status      = resp.GetStatusCode()
	)

	content, err := resp.GetContent()
	if err != nil {
		logger.Error(fmt.Errorf("getting content: %w", err))
	}

	if status == 0 {
		logger.Error(fmt.Errorf("attempted to return status code 0"))
		status = http.StatusInternalServerError
	}

	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	} else {
		w.Header().Set("Content-Length", "0")
	}

	for k, v := range headers {
		if _, ok := w.Header()[k]; !ok {
			w.Header().Set(k, v)
		}
	}

	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	if len(content) > 0 {
		if _, err := w.Write(content); err != nil {
			logger.Error(fmt.Errorf("writing response: %w", err))
		}
	}
}
