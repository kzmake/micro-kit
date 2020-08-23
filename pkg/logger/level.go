package logger

import (
	"github.com/rs/zerolog"
)

// Level defines log levels.
type Level = zerolog.Level

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = zerolog.DebugLevel
	// InfoLevel defines info log level.
	InfoLevel Level = zerolog.InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel Level = zerolog.WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel Level = zerolog.ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel Level = zerolog.FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel Level = zerolog.PanicLevel
	// NoLevel defines an absent log level.
	NoLevel Level = zerolog.NoLevel
	// Disabled disables the logger.
	Disabled Level = zerolog.Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = zerolog.TraceLevel
)

// ParseLevel は指定された文字列から Level を生成します。
func ParseLevel(level string) Level {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	case "panic":
		return PanicLevel
	case "no":
		return NoLevel
	case "disabled":
		return Disabled
	case "trace":
		return TraceLevel
	default:
		return InfoLevel
	}
}
