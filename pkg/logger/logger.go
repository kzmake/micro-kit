package logger

// DefaultLogger ...
var DefaultLogger = New()

// LevelLogger ...
type LevelLogger interface {
	Tracef(ormat string, args ...interface{})
	Debugf(ormat string, args ...interface{})
	Infof(ormat string, args ...interface{})
	Warnf(ormat string, args ...interface{})
	Errorf(ormat string, args ...interface{})
}

// interfaces
var _ LevelLogger = (*Logger)(nil)

// Tracef ...
func Tracef(format string, args ...interface{}) { DefaultLogger.Tracef(format, args...) }

// Debugf ...
func Debugf(format string, args ...interface{}) { DefaultLogger.Debugf(format, args...) }

// Infof ...
func Infof(format string, args ...interface{}) { DefaultLogger.Infof(format, args...) }

// Warnf ...
func Warnf(format string, args ...interface{}) { DefaultLogger.Warnf(format, args...) }

// Errorf ...
func Errorf(format string, args ...interface{}) { DefaultLogger.Errorf(format, args...) }
