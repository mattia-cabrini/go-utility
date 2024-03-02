// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var syncIoGuard = &sync.Mutex{}

// Fprintf but synced
func Fprintf(fp io.Writer, format string, args ...any) {
	defer Monitor(syncIoGuard)()

	fmt.Fprint(fp, format, args)
}

// Printf but synced
func Printf(format string, args ...any) {
	defer Monitor(syncIoGuard)()

	fmt.Fprint(os.Stdout, format, args)
}
