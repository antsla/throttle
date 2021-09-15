package main

import (
	"fmt"
	"os"

	"throttle/internal/config"
	"throttle/internal/datastore"
	"throttle/internal/log"
	"throttle/internal/storage"
	"throttle/internal/transport"
)

func main() {
	logger := log.InitLogger()

	defer func() {
		if r := recover(); r != nil {
			logger.Panic().Msg(fmt.Sprintf("%s", r))
		}
	}()

	cfg, err := config.Load(logger.With().Str("package", "config").Logger())
	if err != nil {
		logger.Error().Err(err).Msg("Terminate execution")
		os.Exit(1)
	}

	redis := datastore.InitRedis(cfg.Redis, logger.With().Str("package", "datastore").Logger())
	redisSrv := storage.BuildRedisService(logger.With().Str("package", "cache").Logger(), redis)

	server := transport.NewServer(cfg.HTTPBind, logger.With().Str("package", "transport").Logger(), redisSrv)
	err = server.Start()
	if err != nil {
		logger.Error().Err(err).Msg("Server hasn't been started.")
		os.Exit(1)
	}
}
