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
	FATAL LogLevel = (1 << iota) >> 1
	ERROR
	WARNING
	INFO
	VERBOSE
)

var levelToString = map[LogLevel]string{
	FATAL:   " FATAL ",
	ERROR:   " ERROR ",
	WARNING: "WARNING",
	INFO:    " INFO  ",
	VERBOSE: "VERBOSE",
}

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

var logGuard *sync.Mutex = &sync.Mutex{}
var MaximumLevel LogLevel = WARNING

func Logf(level LogLevel, format string, args ...interface{}) {
	logGuard.Lock()
	defer logGuard.Unlock()

	if level > MaximumLevel {
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
