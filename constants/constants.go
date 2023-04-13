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
	ErrConnectingToDatabase = errors.New("failed to establish db connection")
	ErrParticipantNotFound  = errors.New("participant not found in dal")
	ErrLeaderboardNotFound  = errors.New("leaderboard not found in dal")
)
