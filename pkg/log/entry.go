package log

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

type Entry struct {
	Logger  *Logger
	Data    Fields
	Time    time.Time
	Level   Level
	Caller  *runtime.Frame
	Message string
	Buffer  *bytes.Buffer
	Context context.Context
	err     string
}

func NewEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
		// Default is three fields, plus one optional, Give a little extra room
		Data: make(Fields, 6),
	}
}

func (e *Entry) Dup() *Entry {
	data := make(Fields, len(e.Data))
	for k, v := range e.Data {
		data[k] = v
	}
	return &Entry{Logger: e.Logger, Data: data, Time: e.Time, Context: e.Context, err: e.err}
}

func (e *Entry) log(level Level, msg string) {
	var buffer *bytes.Buffer

	newEntry := e.Dup()
	if newEntry.Time.IsZero() {
		newEntry.Time = time.Now()
	}
	newEntry.Level = level
	newEntry.Message = msg

	buffer = getBuffer()
	defer func() {
		newEntry.Buffer = nil
		putBuffer(buffer)
	}()
	buffer.Reset()
	newEntry.Buffer = buffer

	newEntry.write()

	newEntry.Buffer = nil

	if level <= PanicLevel {
		panic(newEntry)
	}
}

func (e *Entry) write() {
	serialized, err := e.Logger.Formatter.Format(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		return
	}
	e.Logger.mu.Lock()
	if _, err := e.Logger.Output.Write(serialized); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
	}
	e.Logger.mu.Unlock()
}

func (e *Entry) Log(level Level, args ...interface{}) {
	if e.Logger.IsLevelEnable(level) {
		e.log(level, fmt.Sprint(args...))
	}
}

func (e *Entry) Logf(level Level, format string, args ...interface{}) {
	if e.Logger.IsLevelEnable(level) {
		e.Log(level, fmt.Sprintf(format, args...))
	}
}
