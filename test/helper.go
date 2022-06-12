package test

import (
	"github.com/byyjoww/leaderboard/config"
	"github.com/byyjoww/leaderboard/dal"
)

func GetTestDalFactory() dal.Factory {
	config := GetTestConfig()
	return dal.NewPgFactory(dal.Config{
		User:              config.Postgres.User,
		Pass:              config.Postgres.Pass,
		Host:              config.Postgres.Host,
		Port:              config.Postgres.Port,
		Database:          config.Postgres.Database,
		PoolSize:          config.Postgres.PoolSize,
		MaxRetries:        config.Postgres.MaxRetries,
		ConnectionTimeout: config.Postgres.ConnectionTimeout,
	})
}

func GetTestConfig() config.Config {
	vpr, err := config.ConfigureViper("../../config/config.yaml", "app", "yaml")
	if err != nil {
		panic("failed to create viper instance")
	}

	return config.ParseAll(vpr)
}
