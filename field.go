package logfield

import (
	"fmt"
	"log"
	"strings"
)

var (
	// defaultFieldFormatter default field formatter
	defaultFieldFormatter FieldFormatter = DefaultFieldFormatter
	// defaultFormatter default fields formatter
	defaultFormatter Formatter = DefaultFormatter
)

// FieldFormatter formatter for field
type FieldFormatter func(Field) string

// Formatter formatter for fields
type Formatter func(Fields) string

// Field log field
type Field struct {
	key   string
	value interface{}
}

func (f Field) String() string {
	return defaultFieldFormatter(f)
}

// DefaultFieldFormatter format field using default format
func DefaultFieldFormatter(f Field) string {
	if f.key != "" && f.value != nil {
		return fmt.Sprintf("%s = %s", f.key, f.value)
	} else if f.value == nil {
		return f.key
	} else {
		return fmt.Sprint(f.value)
	}
}

// DefaultFormatter format fields using default format
func DefaultFormatter(f Fields) string {
	var msg strings.Builder
	prefix := ""
	for key, val := range f {
		msg.WriteString(prefix)
		msg.WriteString(defaultFieldFormatter(Field{key, val}))
	}
	return msg.String()
}

// Fields log fields
type Fields map[string]interface{}

func (f Fields) String() string {
	return defaultFormatter(f)
}

// Add key value in field
func (fields Fields) Add(f Field) {
	fields[f.key] = f.value
}

// SetFormatter set default formatter
func SetFieldFormatter(f FieldFormatter) {
	if f != nil {
		defaultFieldFormatter = f
	}
}

// SetFormatter set default formatter
func SetFormatter(f Formatter) {
	if f != nil {
		defaultFormatter = f
	}
}

// LoggerWithFields create standard logger with given fields as prefix
func LoggerWithFields(logger *log.Logger, fields Fields) *log.Logger {
	if logger == nil || len(fields) == 0 {
		return logger
	}
	prefix := defaultFormatter(fields)
	p := logger.Prefix()
	if p != "" {
		prefix = fmt.Sprintf("%s %s", p, prefix)
	}
	return log.New(logger.Writer(), prefix, log.Flags())
}
