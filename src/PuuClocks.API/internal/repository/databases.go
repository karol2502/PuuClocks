package repository

import (
	"context"
	"fmt"
	"puuclocks/internal/infrastructure"
)

type Databases interface {
	RedisDB() Redis
	DB() infrastructure.MySql
}

type databases struct {
	redisDB Redis
	db infrastructure.MySql
}

type DatabasesConfig struct {
	RedisConfig RedisConfig
	MySqlConfig infrastructure.MySqlConfig
}

func NewDatabases(config DatabasesConfig) (Databases, error) {
	ctx := context.Background()

	r := newRedis(config.RedisConfig)
	err := r.Health(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't initalise redis db: %w", err)
	}

	db,err := infrastructure.NewMySql(config.MySqlConfig)
	if err != nil {
		return nil, fmt.Errorf("couldn't initalise mysql db: %w", err)
	}

	return databases{
		redisDB: r,
		db: db,
	}, nil
}

func (d databases) RedisDB() Redis {
	return d.redisDB
}

func (d databases) DB() infrastructure.MySql {
	return d.db
}