package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/rvldodo/boilerplate/config"
	"github.com/rvldodo/boilerplate/lib/log"
)

type Redis struct {
	Client *redis.Client
}

var (
	ErrNotOk = errors.New("Redis not okay")
	ErrNil   = redis.Nil
)

func New() (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        config.Envs.RedisAddress,
		PoolTimeout: time.Duration(config.Envs.RedisTimeout),
		DB:          config.Envs.RedisDB,
		Password:    config.Envs.RedisPassword,
	})

	r := &Redis{Client: client}
	return r, nil
}

func Run(ctx context.Context, r *redis.Client) error {
	err := r.Ping(ctx)
	if err.Val() != "PONG" {
		log.Errorf("Redis failed to connect: %v", err.Err())
		return err.Err()
	}
	log.Info("Redis successfully connected")
	return nil
}
