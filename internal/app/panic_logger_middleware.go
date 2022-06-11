package app

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

type PanicLoggerMiddleware struct {
	logger logrus.FieldLogger
}

func NewPanicLoggerMiddleware(logger logrus.FieldLogger) *PanicLoggerMiddleware {
	return &PanicLoggerMiddleware{
		logger: logger,
	}
}

func (p *PanicLoggerMiddleware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			panicData := recover()
			if panicData != nil {
				stackTrace := debug.Stack()
				stackTraceAsString := string(stackTrace)
				p.logger.WithField("stackTrace", stackTraceAsString).Errorf("panic: ", panicData)
				fmt.Println(stackTraceAsString)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
