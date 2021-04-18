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

	mu        sync.Mutex
	entryPool sync.Pool

	// ExitFunc to exit the application, default to os.Exit()
	ExitFunc exitFunc
}

type exitFunc func(int)

func New() *Logger {
	return &Logger{
		Output:    os.Stderr,
		Level:     InfoLevel,
		Formatter: new(JSONFormatter),
	}
}

func (l *Logger) newEntry() *Entry {
	entry, ok := l.entryPool.Get().(*Entry)
	if ok {
		return entry
	}
	return NewEntry(l)
}

func (l *Logger) releaseEntry(entry *Entry) {
	entry.Data = map[string]interface{}{}
	l.entryPool.Put(entry)
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

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.Logf(PanicLevel, format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logf(FatalLevel, format, args...)
	l.Exit(1)
}

func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	if l.IsLevelEnable(level) {
		entry := l.newEntry()
		entry.Logf(level, format, args...)
		l.releaseEntry(entry)
	}
}

func (l *Logger) IsLevelEnable(level Level) bool {
	return l.level() >= level
}

func (l *Logger) level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.Level)))
}

func (l *Logger) Exit(code int) {
	if l.ExitFunc == nil {
		l.ExitFunc = os.Exit
	}
	l.ExitFunc(code)
}
