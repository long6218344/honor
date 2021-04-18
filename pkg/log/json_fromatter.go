package log

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	FieldKeyMsg   = "msg"
	FieldKeyLevel = "level"
	FieldKeyTime  = "time"
)

type fieldKey string

type FieldMap map[fieldKey]string

type JSONFormatter struct {
	TimestampFromat  string
	DisableTimestamp bool
	// DisableHTMLEscape allows disabling html escaping in output
	DisableHTMLEscape bool
	DataKey           string
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
