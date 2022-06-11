package cmd

import (
	"errors"

	"github.com/byyjoww/leaderboard/config"
	"github.com/byyjoww/leaderboard/constants"
	"github.com/byyjoww/leaderboard/servers/api/http"
	"github.com/sirupsen/logrus"
)

func Execute() {
	config := config.BuildConfig()
	logger := configureLogger(config)
	app := http.New(logger, config)
	app.ListenAndServe()
}

func configureLogger(configs config.Config) logrus.FieldLogger {
	logger := logrus.New()

	level, err := getLogLevel(configs.Logging.Level)
	if err != nil {
		panic("failed to get log level")
	}

	logrus.SetLevel(level)
	logger.SetLevel(level)

	logger.Info("Initiating program")
	return logger
}

func getLogLevel(level string) (logrus.Level, error) {
	switch level {
	case constants.LogLevelPanic:
		return logrus.PanicLevel, nil
	case constants.LogLevelFatal:
		return logrus.FatalLevel, nil
	case constants.LogLevelError:
		return logrus.ErrorLevel, nil
	case constants.LogLevelWarn:
		return logrus.WarnLevel, nil
	case constants.LogLevelInfo:
		return logrus.InfoLevel, nil
	case constants.LogLevelDebug:
		return logrus.DebugLevel, nil
	case constants.LogLevelTrace:
		return logrus.TraceLevel, nil
	default:
		return 0, errors.New("unable to determine logging level")
	}
}
