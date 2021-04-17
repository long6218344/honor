package log

type fieldKey string

type FieldMap map[fieldKey]string

type JSONFormatter struct {
	TimestampFromat  string
	DisableTimestamp bool
	DataKey          string
}

func (f *JSONFormatter) Format(entry *Entry) ([]byte, error) {

}
