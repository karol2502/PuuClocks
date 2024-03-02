package repository

import (
	"context"
	"fmt"
)

type Databases interface {
	RedisDB() Redis
}

type databases struct {
	redisDB Redis
}

type DatabasesConfig struct {
	RedisConfig RedisConfig
}

func NewDatabases(config DatabasesConfig) (Databases, error) {
	fmt.Printf("Config: %v\n", config)

	ctx := context.Background()
	r := newRedis(config.RedisConfig)

	err := r.Health(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't initalise databases: %w", err)
	}


	return databases{
		redisDB: r,
	}, nil
}

func (d databases) RedisDB() Redis {
	return d.redisDB
}