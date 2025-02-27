package log

import (
	"strings"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string { return strings.ToUpper(l.String()) }
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z0700"
}

func Init(level Level, opts ...Option) {
	zerolog.DefaultContextLogger = NewLogger(level, opts...).l
}
