package trace

import (
	"github.com/cdnnow-pro/tracer-go"
	"github.com/rs/zerolog"
)

type Hook struct{}

func (c Hook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	if span := tracer.SpanFromContext(e.GetCtx()); span.IsValid() {
		e.Str("trace_id", span.TraceId())
		e.Str("span_id", span.SpanId())
	}
}
