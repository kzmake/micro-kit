package technical

import (
	"os"
	"time"

	"github.com/kzmake/micro-kit/pkg/logger"
)

const defaultSkipFrameCount = 4

// Logger はデフォルトで提供される logger です。
var Logger = logger.New(
	logger.WithOutput(os.Stdout),
	logger.WithTimeFormat(time.RFC3339Nano),
	logger.WithSkipFrameCount(defaultSkipFrameCount),
)

// WithFields は fields をログに含む logger を新しく生成します。
func WithFields(fields map[string]interface{}) *logger.Logger { return Logger.WithFields(fields) }

// Tracef は TRACE レベルのログを出力します。
func Tracef(format string, args ...interface{}) { Logger.Tracef(format, args...) }

// Trace は TRACE レベルのログを出力します。
func Trace(v ...interface{}) { Logger.Trace(v...) }

// Debugf は DEBUG レベルのログを出力します。
func Debugf(format string, args ...interface{}) { Logger.Debugf(format, args...) }

// Debug は DEBUG レベルのログを出力します。
func Debug(v ...interface{}) { Logger.Debug(v...) }

// Infof は INFO レベルのログを出力します。
func Infof(format string, args ...interface{}) { Logger.Infof(format, args...) }

// Info は INFO レベルのログを出力します。
func Info(v ...interface{}) { Logger.Info(v...) }

// Warnf は WARN レベルのログを出力します。
func Warnf(format string, args ...interface{}) { Logger.Warnf(format, args...) }

// Warn は WARN レベルのログを出力します。
func Warn(v ...interface{}) { Logger.Warn(v...) }

// Errorf は ERROR レベルのログを出力します。
func Errorf(format string, args ...interface{}) { Logger.Errorf(format, args...) }

// Error は ERROR レベルのログを出力します。
func Error(v ...interface{}) { Logger.Error(v...) }
