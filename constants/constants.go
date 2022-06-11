package constants

import "errors"

const (
	EnvironmentProduction string = "production"
	EnvironmentStaging    string = "staging"
)

const (
	LogLevelPanic string = "panic"
	LogLevelFatal string = "fatal"
	LogLevelError string = "error"
	LogLevelWarn  string = "warn"
	LogLevelInfo  string = "info"
	LogLevelDebug string = "debug"
	LogLevelTrace string = "trace"
)

var (
	ErrPlayerNotFound      = errors.New("player not found in dal")
	ErrLeaderboardNotFound = errors.New("leaderboard not found in dal")
)
