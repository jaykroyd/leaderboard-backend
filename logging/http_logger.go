package logging

import (
	"github.com/byyjoww/leaderboard/services/http"
)

var (
	_ http.Logger = (*HttpLogger)(nil)
)

type Fields = map[string]interface{}

type HttpLogger struct {
	base Logger
}

func NewHttpLogger(logger Logger) http.Logger {
	return NewHttpLoggerWithLogger(logger)
}

// NewLoggrWithEntry returns a new Logger implementation based on a provided logrus entry instance
func NewHttpLoggerWithEntry(logger Logger) http.Logger {
	return &HttpLogger{base: logger}
}

// NewLoggrWithLogger returns a new Logger implementation based on a provided logrus instance
func NewHttpLoggerWithLogger(logger Logger) http.Logger {
	return &HttpLogger{base: logger}
}

func (l *HttpLogger) Fatal(format ...interface{}) {
	l.base.Fatal(format...)
}

func (l *HttpLogger) Fatalf(format string, args ...interface{}) {
	l.base.Fatalf(format, args...)
}

func (l *HttpLogger) Fatalln(args ...interface{}) {
	l.base.Fatalln(args...)
}

func (l *HttpLogger) Debug(args ...interface{}) {
	l.base.Debug(args...)
}

func (l *HttpLogger) Debugf(format string, args ...interface{}) {
	l.base.Debugf(format, args...)
}

func (l *HttpLogger) Debugln(args ...interface{}) {
	l.base.Debugln(args...)
}

func (l *HttpLogger) Error(args ...interface{}) {
	l.base.Error(args...)
}

func (l *HttpLogger) Errorf(format string, args ...interface{}) {
	l.base.Errorf(format, args...)
}

func (l *HttpLogger) Errorln(args ...interface{}) {
	l.base.Errorln(args...)
}

func (l *HttpLogger) Info(args ...interface{}) {
	l.base.Info(args...)
}

func (l *HttpLogger) Infof(format string, args ...interface{}) {
	l.base.Infof(format, args...)
}

func (l *HttpLogger) Infoln(args ...interface{}) {
	l.base.Infoln(args...)
}

func (l *HttpLogger) Warn(args ...interface{}) {
	l.base.Warn(args...)
}

func (l *HttpLogger) Warnf(format string, args ...interface{}) {
	l.base.Warnf(format, args...)
}

func (l *HttpLogger) Warnln(args ...interface{}) {
	l.base.Warnln(args...)
}

func (l *HttpLogger) Panic(args ...interface{}) {
	l.base.Panic(args...)
}

func (l *HttpLogger) Panicf(format string, args ...interface{}) {
	l.base.Panicf(format, args...)
}

func (l *HttpLogger) Panicln(args ...interface{}) {
	l.base.Panicln(args...)
}

func (l *HttpLogger) WithFields(fields map[string]interface{}) http.Logger {
	return &HttpLogger{base: l.base.WithFields(fields)}
}

func (l *HttpLogger) WithField(key string, value interface{}) http.Logger {
	return &HttpLogger{base: l.base.WithField(key, value)}
}

func (l *HttpLogger) WithError(err error) http.Logger {
	return &HttpLogger{base: l.base.WithError(err)}
}
