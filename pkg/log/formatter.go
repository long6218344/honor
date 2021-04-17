package log

import (
	"bytes"
	"context"
	"runtime"
	"time"
)

type Formatter interface {
	Format(*Entry) ([]byte, error)
}

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
