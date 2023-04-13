package middleware

import (
	"net/http"
	"runtime/debug"

	lib "github.com/byyjoww/leaderboard/services/http"
	app "github.com/byyjoww/leaderboard/services/http/server"
)

var (
	_ app.Middleware = (*PanicLoggerMiddleware)(nil)
)

type PanicLoggerMiddleware struct {
}

func NewPanicLoggerMiddleware() *PanicLoggerMiddleware {
	return &PanicLoggerMiddleware{}
}

func (p *PanicLoggerMiddleware) ServeNext(logger lib.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			panicData := recover()
			if panicData != nil {
				stackTrace := debug.Stack()
				stackTraceAsString := string(stackTrace)
				logger.WithField("stackTrace", stackTraceAsString).Errorf("panic: %v", panicData)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
