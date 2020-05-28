package logger

import (
	"os"
	"time"

	zerolog "github.com/rs/zerolog"
)

// Logger ...
type Logger struct {
	Zerolog zerolog.Logger
}

// New ...
func New() Logger {
	w := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
	}

	logger := Logger{
		Zerolog: zerolog.New(w).With().Timestamp().Logger(),
	}

	return logger
}

// Tracef ...
func (l Logger) Tracef(format string, args ...interface{}) {
	l.Zerolog.Trace().Msgf(format, args...)
}

// Debugf ...
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Zerolog.Debug().Msgf(format, args...)
}

// Infof ...
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Zerolog.Info().Msgf(format, args...)
}

// Warnf ...
func (l Logger) Warnf(format string, args ...interface{}) {
	l.Zerolog.Warn().Msgf(format, args...)
}

// Errorf ...
func (l Logger) Errorf(format string, args ...interface{}) {
	l.Zerolog.Error().Msgf(format, args...)
}
