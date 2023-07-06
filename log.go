/*
MIT License

Copyright (c) 2023 Mattia Cabrini

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
