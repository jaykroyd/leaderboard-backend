package config

import "time"

type Config struct {
	Environment string   `mapstructure:"environment"`
	Logging     Logging  `mapstructure:"logging"`
	Http        HTTP     `mapstructure:"http"`
	Postgres    Postgres `mapstructure:"postgres"`
}

type Logging struct {
	Level string `mapstructure:"level"`
}

type HTTP struct {
	Address string        `mapstructure:"address"`
	Auth    Authorization `mapstructure:"auth"`
}

type Authorization struct {
	Enabled bool   `mapstructure:"enabled"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"pass"`
}

type Postgres struct {
	User              string        `mapstructure:"user"`
	Pass              string        `mapstructure:"pass"`
	Host              string        `mapstructure:"host"`
	Port              string        `mapstructure:"port"`
	Database          string        `mapstructure:"database"`
	PoolSize          int           `mapstructure:"poolSize"`
	MaxRetries        int           `mapstructure:"maxRetries"`
	ConnectionTimeout time.Duration `mapstructure:"connectionTimeout"`
}
