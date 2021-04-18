package log

import "io"

type Fields map[string]interface{}

var defaultLogger = New()

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

// SetOutput sets the standard logger output.
func SetOutput(output io.Writer) {
	defaultLogger.SetOutput(output)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Panic(args ...interface{}) {

}

func Fatal(args ...interface{}) {

}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
}

func Fatalf(format string, arg ...interface{}) {

}
