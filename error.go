// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"fmt"
	"runtime"
)

func Mypanic(err error) {
	if err == nil {
		return
	}

	Logf(FATAL, "%v", err)
}

func Deferrable(f func() error, pre func(), post func()) {
	if pre != nil {
		pre()
	}
	Mypanic(f())
	if post != nil {
		post()
	}
}

func AppendError(err error) error {
	if err == nil {
		return nil
	}
	var fileName = trace()[1]
	return fmt.Errorf("%s: %v", fileName, err)
}

func trace() (receipts []string) {
	var p = make([]uintptr, 10)
	runtime.Callers(2, p)

	for _, pcx := range p {
		f := runtime.FuncForPC(pcx)
		receipts = append(receipts, f.Name())
	}

	return
}
