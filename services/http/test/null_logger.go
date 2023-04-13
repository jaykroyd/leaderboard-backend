package test

import lib "github.com/byyjoww/leaderboard/services/http"

type NullLogger struct {
}

func NewNullLogger() lib.Logger {
	return &NullLogger{}
}

func (l *NullLogger) Fatal(format ...interface{}) {

}

func (l *NullLogger) Fatalf(format string, args ...interface{}) {

}

func (l *NullLogger) Fatalln(args ...interface{}) {

}

func (l *NullLogger) Debug(args ...interface{}) {

}

func (l *NullLogger) Debugf(format string, args ...interface{}) {

}

func (l *NullLogger) Debugln(args ...interface{}) {

}

func (l *NullLogger) Error(args ...interface{}) {

}

func (l *NullLogger) Errorf(format string, args ...interface{}) {

}

func (l *NullLogger) Errorln(args ...interface{}) {

}

func (l *NullLogger) Info(args ...interface{}) {

}

func (l *NullLogger) Infof(format string, args ...interface{}) {

}

func (l *NullLogger) Infoln(args ...interface{}) {

}

func (l *NullLogger) Warn(args ...interface{}) {

}

func (l *NullLogger) Warnf(format string, args ...interface{}) {

}

func (l *NullLogger) Warnln(args ...interface{}) {

}

func (l *NullLogger) Panic(args ...interface{}) {

}

func (l *NullLogger) Panicf(format string, args ...interface{}) {

}

func (l *NullLogger) Panicln(args ...interface{}) {

}

func (l *NullLogger) WithFields(fields map[string]interface{}) lib.Logger {
	return l
}

func (l *NullLogger) WithField(key string, value interface{}) lib.Logger {
	return l
}

func (l *NullLogger) WithError(err error) lib.Logger {
	return l
}
