package log

import (
	"io"
	"log/syslog"
)

func NewSyslogWriter(level Level, identifier string) (io.Writer, error) {
	writer, err := syslog.New(levelToSyslogPriority(level), identifier)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func NewRsyslogWriter(level Level, network, raddr, identifier string) (io.Writer, error) {
	writer, err := syslog.Dial(network, raddr, levelToSyslogPriority(level), identifier)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func levelToSyslogPriority(level Level) syslog.Priority {
	switch level {
	case TraceLevel, DebugLevel:
		return syslog.LOG_DEBUG
	case InfoLevel:
		return syslog.LOG_INFO
	case WarnLevel:
		return syslog.LOG_WARNING
	case ErrorLevel:
		return syslog.LOG_ERR
	case FatalLevel:
		return syslog.LOG_CRIT
	default:
		return syslog.LOG_INFO
	}
}
