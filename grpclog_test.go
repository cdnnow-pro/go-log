package log_test

import (
	"bytes"
	"testing"

	. "github.com/cdnnow-pro/logger-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/grpclog"
)

func setupLogger() (grpclog.LoggerV2, *bytes.Buffer) {
	var buf = new(bytes.Buffer)
	l := NewGrpcLogger(DebugLevel, WithGrpcTimestamp(), WithGrpcOutput(buf))
	return l, buf
}

func TestGrpcLogger(t *testing.T) {
	t.Parallel()

	t.Run("Info", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Info("info message", "with few", "words")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, InfoLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "info messagewith fewwords", log.Message)
	})

	t.Run("Infoln", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Infoln("info message", "with few", "words")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, InfoLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "info messagewith fewwords", log.Message)
	})

	t.Run("Infof", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Infof("formatted %s message with %q string", "info", "quoted")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, InfoLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, `formatted info message with "quoted" string`, log.Message)
	})

	t.Run("Warning", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Warning("warn message", "with few", "words")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, WarnLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "warn messagewith fewwords", log.Message)
	})

	t.Run("Warningln", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Warningln("warn message", "with few", "words")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, WarnLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "warn messagewith fewwords", log.Message)
	})

	t.Run("Warningf", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Warningf("formatted %s message with %q string", "warn", "quoted")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, WarnLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, `formatted warn message with "quoted" string`, log.Message)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Error("error message", "with few", "words")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, ErrorLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "error messagewith fewwords", log.Message)
	})

	t.Run("Errorln", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Errorln("error message", "with few", "words")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, ErrorLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "error messagewith fewwords", log.Message)
	})

	t.Run("Errorf", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l, buf := setupLogger()

		// Act
		l.Errorf("formatted %s message with %q string", "error", "quoted")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, ErrorLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, `formatted error message with "quoted" string`, log.Message)
	})
}

func TestGrpcLogger_V(t *testing.T) {
	t.Parallel()

	t.Run("true", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l := NewGrpcLogger(InfoLevel)

		// Act
		result := l.V(2)

		// Assert
		assert.True(t, result)
	})

	t.Run("false", func(t *testing.T) {
		t.Parallel()

		// Arrange
		l := NewGrpcLogger(ErrorLevel)

		// Act
		result := l.V(2)

		// Assert
		assert.False(t, result)
	})
}
