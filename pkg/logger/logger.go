package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Logger zerolog.Logger

func InitLogger(serviceName string) {
	Logger = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("services", serviceName).
		Logger()
}

func InfoLog(msg string, fields map[string]interface{}) {
	event := Logger.Info().Str("level", "INFO")
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

func DBLog(query string, duration time.Duration, fields map[string]interface{}) {
	event := Logger.Info().Str("level", "DB").Str("query", query).Dur("duration", duration)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg("Database query executed")
}

func ErrorLog(err error, fields map[string]interface{}) {
	event := Logger.Error().Str("level", "ERROR").Err(err)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg("An error occurred")
}

func FatalLog(err error, fields map[string]interface{}) {
	event := Logger.Fatal().Str("level", "FATAL").Err(err)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg("Fatal error - exiting")
}
