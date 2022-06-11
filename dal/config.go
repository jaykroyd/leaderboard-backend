package dal

import (
	"fmt"
	"time"
)

type Config struct {
	User              string
	Pass              string
	Host              string
	Port              string
	Database          string
	PoolSize          int
	MaxRetries        int
	ConnectionTimeout time.Duration
}

func (c Config) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
