package logger

import (
	"context"
	"io"
	"os"
	"time"
)

// DefaultOptions は デフォルトのオプションを提供します。
var DefaultOptions = Options{
	Out:            os.Stdout,
	Level:          TraceLevel,
	TimeFormat:     time.RFC3339Nano,
	SkipFrameCount: 3,
	Fields:         make(map[string]interface{}),
	Context:        context.Background(),
}

// Option は Options を更新するための func の定義です。
type Option func(*Options)

// Options は logger 生成時に設定するオプションの定義です。
type Options struct {
	Out            io.Writer
	Level          Level
	TimeFormat     string
	SkipFrameCount int
	Fields         map[string]interface{}
	Context        context.Context
}

// NewOptions は opts をもとに Options を生成します。
func NewOptions(opts ...Option) Options {
	options := DefaultOptions

	for _, o := range opts {
		o(&options)
	}

	return options
}

// WithOutput は output 用の writer を設定します。
func WithOutput(out io.Writer) Option {
	return func(args *Options) {
		args.Out = out
	}
}

// WithLevel は logger の level を設定します。
func WithLevel(level Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

// WithTimeFormat は time format を設定します。
func WithTimeFormat(timeFormat string) Option {
	return func(args *Options) {
		args.TimeFormat = timeFormat
	}
}

// WithSkipFrameCount は logger に caller の skip frame count を設定します。
func WithSkipFrameCount(sfc int) Option {
	return func(args *Options) {
		args.SkipFrameCount = sfc
	}
}

// WithFields は logger に fields を設定します。
func WithFields(fields map[string]interface{}) Option {
	return func(args *Options) {
		args.Fields = fields
	}
}
