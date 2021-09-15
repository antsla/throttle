package config

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
)

type Config struct {
	HTTPBind string
	Redis    RedisConf
}

type RedisConf struct {
	Host     string
	Port     string
	Password string
}

func Load(log zerolog.Logger) (Config, error) {
	cfg := Config{}

	httpBind := os.Getenv("HTTP_BIND")
	if len(httpBind) == 0 {
		log.Warn().Msg("env HTTP_BIND is empty, use default value [8080]")
		httpBind = "8080"
	}
	cfg.HTTPBind = httpBind

	redisHost := os.Getenv("REDIS_HOST")
	if len(redisHost) == 0 {
		return cfg, errors.New("env REDIS_HOST is empty")
	}
	cfg.Redis.Host = redisHost

	redisPort := os.Getenv("REDIS_PORT")
	if len(redisPort) == 0 {
		return cfg, errors.New("env REDIS_PORT is empty")
	}
	cfg.Redis.Port = redisPort

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if len(redisPassword) == 0 {
		return cfg, errors.New("env REDIS_PASSWORD is empty")
	}
	cfg.Redis.Password = redisPassword

	return cfg, nil
}
