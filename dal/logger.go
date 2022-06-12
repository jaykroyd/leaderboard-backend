package dal

import (
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

type PgLogger struct {
	logger logrus.FieldLogger
}

func NewPgLogger(logger logrus.FieldLogger) *PgLogger {
	return &PgLogger{
		logger: logger,
	}
}

func (l *PgLogger) BeforeQuery(event *pg.QueryEvent) {

}

func (l *PgLogger) AfterQuery(event *pg.QueryEvent) {
	l.logger.Info(event.FormattedQuery())
}
