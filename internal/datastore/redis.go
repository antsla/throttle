package datastore

import (
	"context"
	"fmt"
	"os"
	"throttle/internal/config"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
)

func InitRedis(cfg config.RedisConf, log zerolog.Logger) *redis.Client {
	redisStorage := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
	})

	_, err := redisStorage.Ping(context.Background()).Result()
	if err != nil {
		log.Error().Err(err).Msg("couldn't connect to redis")
		os.Exit(1)
	}

	return redisStorage
}
