package log

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	FieldKeyMsg           = "msg"
	FieldKeyLevel         = "level"
	FieldKeyTime          = "time"
	FieldKeyLogHonorError = "honor_error"
	FieldKeyFunc          = "func"
	FieldKeyFile          = "file"

	defaultTimestampFormat = "2006-01-02 15:04:05.000"
)

type fieldKey string

type FieldMap map[fieldKey]string

func (f FieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}
	return string(key)
}

type jsonFormatter struct{}

func (f *jsonFormatter) Format(entry *Entry) ([]byte, error) {
	data := make(Fields, 4)
	if entry.err != "" {
		data[FieldKeyLogHonorError] = entry.err
	}
	data[FieldKeyTime] = entry.Time.Format(defaultTimestampFormat)
	data[FieldKeyMsg] = entry.Message
	data[FieldKeyLevel] = entry.Level.String()

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return b.Bytes(), nil
}
