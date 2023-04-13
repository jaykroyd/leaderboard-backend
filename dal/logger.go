package dal

import (
	"github.com/byyjoww/leaderboard/logging"
	"github.com/go-pg/pg"
)

type PgLogger struct {
	logger logging.Logger
}

func NewPgLogger(logger logging.Logger) *PgLogger {
	return &PgLogger{
		logger: logger,
	}
}

func (l *PgLogger) BeforeQuery(event *pg.QueryEvent) {

}

func (l *PgLogger) AfterQuery(event *pg.QueryEvent) {
	l.logger.Info(event.FormattedQuery())
}
