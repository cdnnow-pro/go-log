// SPDX-License-Identifier: MIT

package log

import (
	"context"
	"os"
	"sync/atomic"

	"github.com/rs/zerolog"
)

var (
	callerEnabled        atomic.Bool
	deduplicationEnabled atomic.Bool
)

type Logger struct {
	l *zerolog.Logger
}

// SetGlobalLevel creates a logger with specified level and stores it as default logger.
func SetGlobalLevel(level Level) {
	l := zerolog.DefaultContextLogger.Level(zerolog.Level(level))
	zerolog.DefaultContextLogger = &l
}

// SetCallerEnabled sets the global flag that determines whether to add in information
// about the log point ("file:line") in the log entry.
func SetCallerEnabled(enabled bool) {
	callerEnabled.Store(enabled)
}

// SetDeduplicationEnabled sets the global flag that determines whether to uniqualize
// custom fields keys.
//
// Must be called before using of the custom fields.
func SetDeduplicationEnabled(enabled bool) {
	deduplicationEnabled.Store(enabled)
}

func NewLogger(level Level, opts ...Option) *Logger {
	zl := zerolog.New(os.Stdout).
		Level(zerolog.Level(level))

	l := &Logger{l: &zl}
	for _, opt := range opts {
		opt(l)
	}

	return l
}

func ToContext(ctx context.Context, logger *Logger) context.Context {
	return logger.l.WithContext(ctx)
}

func FromContext(ctx context.Context) *Logger {
	return &Logger{l: zerolog.Ctx(ctx)}
}

func (l *Logger) GetLevel() Level {
	return Level(l.l.GetLevel())
}

func (l *Logger) Level(level Level) *Logger {
	zl := l.l.Level(zerolog.Level(level))
	return &Logger{l: &zl}
}

func (l *Logger) DebugWithTrace(ctx context.Context, msg, trace string, fields ...any) {
	event := l.l.Debug()
	event = withFieldsAndCaller(ctx, event, fields)
	if l.GetLevel() == TraceLevel {
		event = event.Str("trace", trace)
	}
	event.Msg(msg)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...any) {
	event := l.l.Debug()
	event = withFieldsAndCaller(ctx, event, fields)
	event.Msg(msg)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...any) {
	event := l.l.Info()
	event = withFieldsAndCaller(ctx, event, fields)
	event.Msg(msg)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...any) {
	event := l.l.Warn()
	event = withFieldsAndCaller(ctx, event, fields)
	event.Msg(msg)
}

func (l *Logger) Error(ctx context.Context, err error, msg string, fields ...any) {
	event := l.l.Error().Err(err)
	event = withFieldsAndCaller(ctx, event, fields)
	event.Msg(msg)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...any) {
	event := l.l.Fatal()
	event = withFieldsAndCaller(ctx, event, fields)
	event.Msg(msg)
}

func (l *Logger) FatalError(ctx context.Context, err error, msg string, fields ...any) {
	event := l.l.Fatal().Err(err)
	event = withFieldsAndCaller(ctx, event, fields)
	event.Msg(msg)
}

func (l *Logger) Force(ctx context.Context, msg string, fields ...any) {
	l2 := l.l.Level(zerolog.InfoLevel)
	event := l2.Info()
	event = withFieldsAndCaller(ctx, event, fields)
	event.Msg(msg)
}

func DebugWithTrace(ctx context.Context, msg, trace string, fields ...any) {
	FromContext(ctx).DebugWithTrace(ctx, msg, trace, fields...)
}

func Debug(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Debug(ctx, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Info(ctx, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Warn(ctx, msg, fields...)
}

func Error(ctx context.Context, err error, msg string, fields ...any) {
	FromContext(ctx).Error(ctx, err, msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Fatal(ctx, msg, fields...)
}

func FatalError(ctx context.Context, err error, msg string, fields ...any) {
	FromContext(ctx).FatalError(ctx, err, msg, fields...)
}

func Force(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Force(ctx, msg, fields...)
}

func withFieldsAndCaller(ctx context.Context, event *zerolog.Event, f Fields) *zerolog.Event {
	if callerEnabled.Load() {
		event = event.Caller(3) //nolint:mnd
	}
	return event.Fields([]any(ExtractFields(ctx).With(f)))
}
