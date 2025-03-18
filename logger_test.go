// SPDX-License-Identifier: MIT

package log_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"testing"

	. "github.com/cdnnow-pro/logger-go"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type TestLogType struct {
	WithField1 any `json:"with_field_1,omitempty"`
	Trace      any `json:"trace,omitempty"`

	Level   string `json:"level"`
	Message string `json:"msg"`

	Time   string `json:"time,omitempty"`
	Caller string `json:"caller,omitempty"`
	Error  string `json:"error,omitempty"`
}

func getTestData() (context.Context, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	l := NewLogger(DebugLevel, WithTimestamp(), WithOutput(buf))
	return ToContext(context.Background(), l), buf
}

func getTestDataTextPlain() (context.Context, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	l := NewLogger(DebugLevel, WithPlainText(false, buf))
	return ToContext(context.Background(), l), buf
}

func newTestLogType(data []byte) TestLogType {
	log := TestLogType{}
	_ = json.Unmarshal(data, &log)
	return log
}

func Test_JsonLog(t *testing.T) {
	t.Parallel()

	t.Run("debug", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Debug(ctx, "debugMsg")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, DebugLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "debugMsg", log.Message)
	})

	t.Run("debug with KV", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Debug(ctx, "debugMsg", "with_field_1", "test")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, DebugLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "debugMsg", log.Message)
		assert.Equal(t, "test", log.WithField1)
	})

	t.Run("debug with trace and KV", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()
		ctx = ToContext(ctx, FromContext(ctx).Level(TraceLevel))

		// Act
		DebugWithTrace(ctx, "debugMsg", "traceData", "with_field_1", "test")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, DebugLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "debugMsg", log.Message)
		assert.Equal(t, "test", log.WithField1)
		assert.Equal(t, "traceData", log.Trace)
	})

	t.Run("info", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Info(ctx, "infoMsg")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, InfoLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "infoMsg", log.Message)
	})

	t.Run("info with KV", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Info(ctx, "infoMsg", "with_field_1", "test")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, InfoLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "infoMsg", log.Message)
		assert.Equal(t, "test", log.WithField1)
	})

	t.Run("warn", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Warn(ctx, "warnMsg")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, WarnLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "warnMsg", log.Message)
	})

	t.Run("warn with KV", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Warn(ctx, "warnMsg", "with_field_1", "test")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, WarnLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "warnMsg", log.Message)
		assert.Equal(t, "test", log.WithField1)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Error(ctx, errors.New("test error"), "errorMsg")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, ErrorLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "errorMsg", log.Message)
		assert.Equal(t, "test error", log.Error)
	})

	t.Run("force", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Force(ctx, "forceMsg")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, InfoLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "forceMsg", log.Message)
	})

	t.Run("force with KV", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestData()

		// Act
		Force(ctx, "forceMsg", "with_field_1", "test")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, InfoLevel.String(), log.Level)
		assert.NotEmpty(t, log.Time)
		assert.Equal(t, "forceMsg", log.Message)
		assert.Equal(t, "test", log.WithField1)
	})
}

func Test_PlainTextLog(t *testing.T) {
	t.Parallel()

	t.Run("debug", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestDataTextPlain()

		// Act
		Debug(ctx, "debugMsg")

		// Assert
		assert.Contains(t, buf.String(), "DBG")
		assert.Contains(t, buf.String(), "debugMsg")
	})

	t.Run("debug with trace", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestDataTextPlain()
		ctx = ToContext(ctx, FromContext(ctx).Level(TraceLevel))

		// Act
		DebugWithTrace(ctx, "debugMsg", "traceData")

		// Assert
		assert.Contains(t, buf.String(), "DBG")
		assert.Contains(t, buf.String(), "debugMsg")
		assert.Contains(t, buf.String(), "traceData")
	})

	t.Run("info", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestDataTextPlain()

		// Act
		Info(ctx, "infoMsg")

		// Assert
		assert.Contains(t, buf.String(), "INF")
		assert.Contains(t, buf.String(), "infoMsg")
	})

	t.Run("warn", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestDataTextPlain()

		// Act
		Warn(ctx, "warnMsg")

		// Assert
		assert.Contains(t, buf.String(), "WRN")
		assert.Contains(t, buf.String(), "warnMsg")
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestDataTextPlain()

		// Act
		Error(ctx, errors.New("test error"), "errorMsg")

		// Assert
		assert.Contains(t, buf.String(), "ERR")
		assert.Contains(t, buf.String(), "errorMsg")
		assert.Contains(t, buf.String(), `test error`)
	})

	t.Run("force", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx, buf := getTestDataTextPlain()

		// Act
		Force(ctx, "forceMsg")

		// Assert
		s := buf.String()
		assert.Contains(t, s, "forceMsg")
	})
}

func TestLogger_Level(t *testing.T) {
	t.Parallel()

	// Arrange
	ctx, buf := getTestData()
	l := FromContext(ctx).Level(InfoLevel)

	// Act
	l.Debug(ctx, "hidden")
	l.Info(ctx, "visible")

	// Assert
	log := buf.String()
	assert.Contains(t, log, "visible")
	assert.NotContains(t, log, "hidden")
}

func TestLogger_Force(t *testing.T) {
	t.Parallel()

	// Arrange
	ctx, buf := getTestData()
	l := FromContext(ctx).Level(InfoLevel)

	// Act
	l.Debug(ctx, "hidden")
	l.Info(ctx, "visible")
	l.Force(ctx, "forced")

	// Assert
	log := buf.String()
	assert.Contains(t, log, "visible")
	assert.NotContains(t, log, "hidden")
	assert.Contains(t, log, "forced")
}

func Test_Caller(t *testing.T) {
	t.Parallel()

	// Arrange
	ctx, buf := getTestData()

	t.Run("enabled", func(t *testing.T) {
		// Act
		SetCallerEnabled(true)
		Info(ctx, "infoMsg")
		SetCallerEnabled(false)

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Contains(t, log.Caller, `/logger_test.go:`)
		buf.Reset()
	})

	t.Run("disabled", func(t *testing.T) {
		t.Parallel()

		// Act
		Info(ctx, "infoMsg")

		// Assert
		log := newTestLogType(buf.Bytes())
		assert.Equal(t, log.Caller, ``)
	})
}

func TestParseLevel(t *testing.T) {
	t.Parallel()

	t.Run("invalid level", func(t *testing.T) {
		t.Parallel()

		// Act
		level, err := ParseLevel("invalid")

		// Assert
		assert.Equal(t, ErrorLevel, level)
		assert.Error(t, err)
	})

	t.Run("max level", func(t *testing.T) {
		t.Parallel()

		// Act
		level, err := ParseLevel(zerolog.FatalLevel.String())

		// Assert
		assert.Equal(t, Level(zerolog.ErrorLevel), level)
		assert.NoError(t, err)
	})

	t.Run("valid levels", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			levelStr string
			level    Level
		}{
			{"trace", TraceLevel},
			{"debug", DebugLevel},
			{"info", InfoLevel},
			{"warn", WarnLevel},
			{"error", ErrorLevel},
		}
		for _, tt := range tests {
			// Act
			level, err := ParseLevel(tt.levelStr)

			// Assert
			assert.Equal(t, tt.level, level)
			assert.NoError(t, err)
		}
	})
}
