// SPDX-License-Identifier: MIT

package log

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

type Level int8

const (
	TraceLevel = Level(zerolog.TraceLevel)
	DebugLevel = Level(zerolog.DebugLevel)
	InfoLevel  = Level(zerolog.InfoLevel)
	WarnLevel  = Level(zerolog.WarnLevel)
	ErrorLevel = Level(zerolog.ErrorLevel)
	FatalLevel = Level(zerolog.FatalLevel)

	minAllowedLevel = TraceLevel
	maxAllowedLevel = Level(zerolog.ErrorLevel)
)

func (l Level) String() string {
	return strings.ToUpper(zerolog.Level(l).String())
}

// ParseLevel returns Level according specified value (case-insensitive).
//
// If invalid value specified, then an error returned with default level (ErrorLevel).
func ParseLevel(val string) (Level, error) {
	zLevel, err := zerolog.ParseLevel(val)
	if err != nil {
		errMessage := strings.Replace(err.Error(), "defaulting to NoLevel", "defaulting to ErrorLevel", 1)
		return ErrorLevel, fmt.Errorf("cannot parse log level %q: %s", val, errMessage)
	}

	switch level := Level(zLevel); {
	case level < minAllowedLevel:
		fmt.Printf("log level %q less than min allowed\n", val)
		return minAllowedLevel, nil
	case level > maxAllowedLevel:
		fmt.Printf("log level %q greater than max allowed\n", val)
		return maxAllowedLevel, nil
	default:
		return level, nil
	}
}
