// SPDX-License-Identifier: MIT

package metrics

import (
	"github.com/cdnnow-pro/go-metrics"
	"github.com/rs/zerolog"
)

var (
	_counters = metrics.NewCounterVec("log_messages_total", "Number of logged messages", []string{"level"})

	traceCounter = _counters.WithLabelValues("trace")
	debugCounter = _counters.WithLabelValues("debug")
	infoCounter  = _counters.WithLabelValues("info")
	warnCounter  = _counters.WithLabelValues("warn")
	errorCounter = _counters.WithLabelValues("error")
	fatalCounter = _counters.WithLabelValues("fatal")
)

type Hook struct{}

func (c Hook) Run(_ *zerolog.Event, level zerolog.Level, _ string) {
	switch level {
	case zerolog.TraceLevel:
		traceCounter.Inc()
	case zerolog.DebugLevel:
		debugCounter.Inc()
	case zerolog.InfoLevel:
		infoCounter.Inc()
	case zerolog.WarnLevel:
		warnCounter.Inc()
	case zerolog.ErrorLevel:
		errorCounter.Inc()
	case zerolog.FatalLevel:
		fatalCounter.Inc()
	}
}
