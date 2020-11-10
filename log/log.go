package log

import (
	"io"

	"github.com/rs/zerolog"
	"github.com/souleb/buildeploy/app"
)

// Ensure Logger implements app.LoggerService.
var _ app.LoggerService = (*Logger)(nil)

type Logger struct {
	logger *zerolog.Logger
}

type Options func(logger *Logger)

func New(w io.Writer, level string) *Logger {
	var logger zerolog.Logger
	switch level {
	case "info":
		logger = zerolog.New(w).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	case "debug":
		logger = zerolog.New(w).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	}

	l := Logger{
		logger: &logger,
	}

	return &l
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Fatal(err error) {
	l.logger.Fatal().Err(err)
}
