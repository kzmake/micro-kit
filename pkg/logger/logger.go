package logger

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

// Logger は logger の定義です。
type Logger struct {
	*zerolog.Logger
	opts []Option
}

// New は logger を生成します。
func New(opts ...Option) *Logger {
	options := NewOptions(opts...)

	l := &Logger{opts: opts}
	l.applyOptions(options)

	return l
}

// Clone は logger を複製します。
func (l Logger) Clone(v ...interface{}) *Logger {
	return New(l.opts...)
}

// WithFields は fields をログに含む logger を複製します。
func (l Logger) WithFields(fields map[string]interface{}) *Logger {
	clone := l.Clone()
	clone.applyZerolog(l.Logger.With().Fields(fields).Logger())

	return clone
}

// Tracef は TRACE レベルのログを出力します。
func (l Logger) Tracef(format string, args ...interface{}) {
	l.Logger.Trace().Msgf(format, args...)
}

// Trace は TRACE レベルのログを出力します。
func (l Logger) Trace(v ...interface{}) {
	l.Logger.Trace().Msg(fmt.Sprint(v...))
}

// Debugf は DEBUG レベルのログを出力します。
func (l Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debug().Msgf(format, args...)
}

// Debug は DEBUG レベルのログを出力します。
func (l Logger) Debug(v ...interface{}) {
	l.Logger.Debug().Msg(fmt.Sprint(v...))
}

// Infof は INFO レベルのログを出力します。
func (l Logger) Infof(format string, args ...interface{}) {
	l.Logger.Info().Msgf(format, args...)
}

// Info は INFO レベルのログを出力します。
func (l Logger) Info(v ...interface{}) {
	l.Logger.Info().Msg(fmt.Sprint(v...))
}

// Warnf は WARN レベルのログを出力します。
func (l Logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warn().Msgf(format, args...)
}

// Warn は WARN レベルのログを出力します。
func (l Logger) Warn(v ...interface{}) {
	l.Logger.Warn().Msg(fmt.Sprint(v...))
}

// Errorf は ERROR レベルのログを出力します。
func (l Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Error().Msgf(format, args...)
}

// Error は ERROR レベルのログを出力します。
func (l Logger) Error(v ...interface{}) {
	l.Logger.Error().Msg(fmt.Sprint(v...))
}

func (l *Logger) applyOptions(options Options) {
	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "time"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = options.TimeFormat

	log := zerolog.New(options.Out).Level(options.Level).
		With().Timestamp().Stack().CallerWithSkipFrameCount(options.SkipFrameCount).Logger()

	if options.Fields != nil {
		log = log.With().Fields(options.Fields).Logger()
	}

	l.Logger = &log
}

func (l *Logger) applyZerolog(log zerolog.Logger) { // nolint:gocritic
	l.Logger = &log
}
