package storage

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
)

const ttl = 60 // minutes

type RedisService struct {
	log   zerolog.Logger
	redis *redis.Client
}

func BuildRedisService(log zerolog.Logger, redis *redis.Client) RedisService {
	return RedisService{
		log:   log,
		redis: redis,
	}
}

func (rs RedisService) Increment(ctx context.Context, key string) error {
	val, err := rs.redis.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			rs.log.Error().Err(err).Msg("cache error has occurred")
			return err
		}
		val = "0"
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		rs.log.Error().Err(err).Msg("convert error has occurred")
		return err
	}

	intVal++
	err = rs.redis.Set(ctx, key, intVal, time.Duration(ttl)*time.Minute).Err()
	if err != nil {
		rs.log.Error().Err(err).Msg("couldn't write cache")
		return err
	}

	return nil
}

// GetSession - get cache session data
func (rs RedisService) Get(ctx context.Context, key string) (int, error) {
	val, err := rs.redis.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			rs.log.Error().Err(err).Msg("cache error has occurred")
			return 0, err
		}
		return 0, nil
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		rs.log.Error().Err(err).Msg("cache error has occurred")
		return 0, err
	}

	return intVal, nil
}
