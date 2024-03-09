package repository

import (
	"context"
	"fmt"
	"puuclocks/internal/infrastructure"
)

type Databases interface {
	RedisDB() Redis
	DB() infrastructure.MySQL
}

type databases struct {
	redisDB Redis
	db      infrastructure.MySQL
}

type DatabasesConfig struct {
	RedisConfig RedisConfig
	MySQLConfig infrastructure.MySQLConfig
}

func NewDatabases(config *DatabasesConfig) (Databases, error) {
	ctx := context.Background()

	r := newRedis(config.RedisConfig)
	err := r.Health(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't initialize redis db: %w", err)
	}

	db, err := infrastructure.NewMySQL(config.MySQLConfig)
	if err != nil {
		return nil, fmt.Errorf("couldn't initialize mysql db: %w", err)
	}

	return databases{
		redisDB: r,
		db:      db,
	}, nil
}

func (d databases) RedisDB() Redis {
	return d.redisDB
}

func (d databases) DB() infrastructure.MySQL {
	return d.db
}
