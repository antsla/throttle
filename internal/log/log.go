package log

import (
	"os"

	"github.com/rs/zerolog"
)

func InitLogger() *zerolog.Logger {
	zerolog.TimestampFieldName = "date"
	zerolog.TimeFieldFormat = "2006.01.02 15:04:05"

	logger := zerolog.Logger{}.Output(os.Stdout).With().Timestamp().Logger()

	return &logger
}
