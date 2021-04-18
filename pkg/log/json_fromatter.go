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

type JSONFormatter struct {
	TimestampFromat  string
	DisableTimestamp bool
	// DisableHTMLEscape allows disabling html escaping in output
	DisableHTMLEscape bool
	DataKey           string

	// FieldMap allows users to customize the names of keys for default fields.
	FieldMap FieldMap

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}

func (f *JSONFormatter) Format(entry *Entry) ([]byte, error) {
	data := make(Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	if f.DataKey != "" {
		newData := make(Fields, 4)
		newData[f.DataKey] = data
		data = newData
	}

	timestampFormat := f.TimestampFromat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	if entry.err != "" {
		data[f.FieldMap.resolve(FieldKeyLogHonorError)] = entry.err
	}
	if !f.DisableTimestamp {
		data[f.FieldMap.resolve(FieldKeyTime)] = entry.Time.Format(timestampFormat)
	}
	data[f.FieldMap.resolve(FieldKeyMsg)] = entry.Message
	data[f.FieldMap.resolve(FieldKeyLevel)] = entry.Level.String()

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(!f.DisableHTMLEscape)
	if f.PrettyPrint {
		encoder.SetIndent("", " ")
	}
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return b.Bytes(), nil
}
