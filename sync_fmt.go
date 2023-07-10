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

func Fprintf(fp io.Writer, format string, args ...any) {
	defer Monitor(syncIoGuard)()

	fmt.Fprint(fp, format, args)
}

func Printf(format string, args ...any) {
	defer Monitor(syncIoGuard)()

	fmt.Fprint(os.Stdout, format, args)
}
