package logging

import (
	"github.com/byyjoww/leaderboard/config"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	base logrus.FieldLogger
}

// New returns a new Logger implementation based on logrus
func NewLoggr(cfg config.Logging) Logger {
	logger := newLogrus(cfg)
	return NewLoggrWithLogger(logger)
}

// NewLoggrWithEntry returns a new Logger implementation based on a provided logrus entry instance
func NewLoggrWithEntry(logger *logrus.Entry) Logger {
	return &LogrusLogger{base: logger}
}

// NewLoggrWithLogger returns a new Logger implementation based on a provided logrus instance
func NewLoggrWithLogger(logger logrus.FieldLogger) Logger {
	return &LogrusLogger{base: logger}
}

func (l *LogrusLogger) Fatal(format ...interface{}) {
	l.base.Fatal(format...)
}

func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.base.Fatalf(format, args...)
}

func (l *LogrusLogger) Fatalln(args ...interface{}) {
	l.base.Fatalln(args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.base.Debug(args...)
}

func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.base.Debugf(format, args...)
}

func (l *LogrusLogger) Debugln(args ...interface{}) {
	l.base.Debugln(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.base.Error(args...)
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.base.Errorf(format, args...)
}

func (l *LogrusLogger) Errorln(args ...interface{}) {
	l.base.Errorln(args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.base.Info(args...)
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.base.Infof(format, args...)
}

func (l *LogrusLogger) Infoln(args ...interface{}) {
	l.base.Infoln(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.base.Warn(args...)
}

func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.base.Warnf(format, args...)
}

func (l *LogrusLogger) Warnln(args ...interface{}) {
	l.base.Warnln(args...)
}

func (l *LogrusLogger) Panic(args ...interface{}) {
	l.base.Panic(args...)
}

func (l *LogrusLogger) Panicf(format string, args ...interface{}) {
	l.base.Panicf(format, args...)
}

func (l *LogrusLogger) Panicln(args ...interface{}) {
	l.base.Panicln(args...)
}

func (l *LogrusLogger) WithFields(fields map[string]interface{}) Logger {
	return &LogrusLogger{base: l.base.WithFields(fields)}
}

func (l *LogrusLogger) WithField(key string, value interface{}) Logger {
	return &LogrusLogger{base: l.base.WithField(key, value)}
}

func (l *LogrusLogger) WithError(err error) Logger {
	return &LogrusLogger{base: l.base.WithError(err)}
}

func addField(logger *logrus.Entry, key string, value interface{}) *logrus.Entry {
	logger.Data[key] = value
	return logger
}

func addFields(logger *logrus.Entry, fields map[string]interface{}) *logrus.Entry {
	for k, v := range fields {
		addField(logger, k, v)
	}
	return logger
}
