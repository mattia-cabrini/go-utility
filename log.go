// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

type LogLevel int

const (
	FATAL   LogLevel = (1 << iota) >> 1 // To log an error and exit
	ERROR                               // To log and error
	WARNING                             // To log a warning
	INFO                                // To log some information
	VERBOSE                             // To log some detailed information (meant to debug)
)

// Maps log levels to strings
var levelToString = map[LogLevel]string{
	FATAL:   " FATAL ",
	ERROR:   " ERROR ",
	WARNING: "WARNING",
	INFO:    " INFO  ",
	VERBOSE: "VERBOSE",
}

// Maps log levels to colours
func getLevelColor(level LogLevel) string {
	switch level {
	case FATAL:
		return Red
	case ERROR:
		return Red
	case WARNING:
		return Blue
	case INFO:
		return Green
	}

	return Gray
}

// To log synchronously
var logGuard *sync.Mutex = &sync.Mutex{}

// Minimum level to log
// If it is WARNING, INFO and VERBOSE are not logged
// If it is ERROR, WARNING, INFO and WARNING are not logged
// ... and so on ...
var MinimumLevel LogLevel = WARNING

// Log a message with a specific log level; Format and args are Printf-like
func Logf(level LogLevel, format string, args ...interface{}) {
	logGuard.Lock()
	defer logGuard.Unlock()

	if level > MinimumLevel {
		return
	}

	format = logfCompose(level, format)

	if level == FATAL {
		_, _ = fmt.Fprintf(os.Stderr, format, args...)
		os.Exit(1)
	} else {
		_, _ = fmt.Fprintf(os.Stderr, format, args...)
	}
}

func logfCompose(level LogLevel, format string) string {
	format = "[" +
		getLevelColor(level) + levelToString[level] + Reset +
		"] " + format

	if format[len(format)-1] != '\n' {
		format = format + "\n"
	}
	return format
}

// Log a message on *testing.T with a specific log level; Format and args are Printf-like
func Tlogf(t *testing.T, level LogLevel, format string, args ...interface{}) {
	logGuard.Lock()
	defer logGuard.Unlock()

	format = logfCompose(level, format)

	if level == FATAL || level == ERROR {
		t.Errorf(format, args...)
	} else {
		t.Logf(format, args...)
	}
}
