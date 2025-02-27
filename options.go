package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Option func(*Logger)

func WithHook(h zerolog.Hook) Option {
	return func(l *Logger) {
		zl := l.l.Hook(h)
		l.l = &zl
	}
}

func WithTimestamp() Option {
	return func(l *Logger) {
		zl := l.l.With().Timestamp().Logger()
		l.l = &zl
	}
}

func WithOutput(w io.Writer) Option {
	return func(l *Logger) {
		zl := l.l.Output(w).With().Logger()
		l.l = &zl
	}
}

// WithPlainText creates an output writer with the plain text format instead of JSON.
//
// Optionally the inner output writer could be specified as second argument.
// Otherwise, the os.Stdout will be used.
func WithPlainText(withTimestamp bool, w ...io.Writer) Option {
	if len(w) == 0 {
		w = []io.Writer{os.Stdout}
	}

	partsOrder := []string{
		zerolog.LevelFieldName,
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
	}
	if withTimestamp {
		partsOrder = append([]string{zerolog.TimestampFieldName}, partsOrder...)
	}

	writer := zerolog.NewConsoleWriter(func(writer *zerolog.ConsoleWriter) {
		writer.PartsOrder = partsOrder
		writer.Out = w[0]
	})

	return WithOutput(writer)
}

type GrpcOption func(*grpcLogger)

func WithGrpcTimestamp() GrpcOption {
	return func(l *grpcLogger) {
		zl := l.l.With().Timestamp().Logger()
		l.l = &zl
	}
}

func WithGrpcOutput(w io.Writer) GrpcOption {
	return func(l *grpcLogger) {
		zl := l.l.Output(w).With().Logger()
		l.l = &zl
	}
}

// WithGrpcPlainText creates an output writer with the plain text format instead of JSON.
//
// Optionally the inner output writer could be specified as second argument.
// Otherwise, the os.Stdout will be used.
func WithGrpcPlainText(withTimestamp bool, w ...io.Writer) GrpcOption {
	if len(w) == 0 {
		w = []io.Writer{os.Stdout}
	}

	partsOrder := []string{
		zerolog.LevelFieldName,
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
	}
	if withTimestamp {
		partsOrder = append([]string{zerolog.TimestampFieldName}, partsOrder...)
	}

	writer := zerolog.NewConsoleWriter(func(writer *zerolog.ConsoleWriter) {
		writer.PartsOrder = partsOrder
		writer.Out = w[0]
	})

	return WithGrpcOutput(writer)
}
