// SPDX-License-Identifier: MIT

package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/grpclog"
)

type grpcLogger struct {
	l *zerolog.Logger
}

func NewGrpcLogger(level Level, opts ...GrpcOption) grpclog.LoggerV2 {
	zl := zerolog.New(os.Stdout)
	zl = zl.Level(zerolog.Level(level))

	l := &grpcLogger{l: &zl}
	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (l *grpcLogger) Info(args ...any) {
	l.l.Info().Msg(fmt.Sprint(args...))
}

func (l *grpcLogger) Infof(format string, args ...any) {
	l.l.Info().Msg(fmt.Sprintf(format, args...))
}

func (l *grpcLogger) Infoln(args ...any) {
	l.Info(args...)
}

func (l *grpcLogger) Warning(args ...any) {
	l.l.Warn().Msg(fmt.Sprint(args...))
}

func (l *grpcLogger) Warningf(format string, args ...any) {
	l.l.Warn().Msg(fmt.Sprintf(format, args...))
}

func (l *grpcLogger) Warningln(args ...any) {
	l.Warning(args...)
}

func (l *grpcLogger) Error(args ...any) {
	l.l.Error().Msg(fmt.Sprint(args...))
}

func (l *grpcLogger) Errorf(format string, args ...any) {
	l.l.Error().Msg(fmt.Sprintf(format, args...))
}

func (l *grpcLogger) Errorln(args ...any) {
	l.Error(args...)
}

func (l *grpcLogger) Fatal(args ...any) {
	l.l.Fatal().Msg(fmt.Sprint(args...))
}

func (l *grpcLogger) Fatalf(format string, args ...any) {
	l.l.Fatal().Msg(fmt.Sprintf(format, args...))
}

func (l *grpcLogger) Fatalln(args ...any) {
	l.Fatal(args...)
}

func (l *grpcLogger) V(lvl int) bool {
	return int(l.l.GetLevel()) <= lvl
}
