package log

import (
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type Logger struct {
	Output    io.Writer
	Level     Level
	Formatter Formatter

	mu sync.Mutex
}

func New() *Logger {
	return &Logger{
		Output: os.Stdout,
		Level:  InfoLevel,
	}
}

func (l *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&l.Level), uint32(level))
}

func (l *Logger) SetOutput(output io.Writer) {
	l.mu.Lock()
	l.Output = output
	l.mu.Unlock()
}

func (l *Logger) Info(args ...interface{}) {

}

func (l *Logger) Warn(args ...interface{}) {

}

func (l *Logger) Error(args ...interface{}) {

}

func (l *Logger) Infof(format string, args ...interface{}) {

}

func (l *Logger) Warnf(format string, args ...interface{}) {

}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logf(ErrorLevel, format, args...)
}

func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	if l.IsLevelEnable(level) {

	}
}

func (l *Logger) IsLevelEnable(level Level) bool {
	return l.level() >= level
}

func (l *Logger) level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.Level)))
}
