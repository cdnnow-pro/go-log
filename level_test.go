package log_test

import (
	"testing"

	. "github.com/cdnnow-pro/logger-go"
	"github.com/stretchr/testify/assert"
)

func TestLevel_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		level Level
		want  string
	}{
		{"debug", DebugLevel, "DEBUG"},
		{"info", InfoLevel, "INFO"},
		{"warn", WarnLevel, "WARN"},
		{"error", ErrorLevel, "ERROR"},
		{"fatal", FatalLevel, "FATAL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			str := tt.level.String()

			// Assert
			assert.Equal(t, tt.want, str)
		})
	}
}
