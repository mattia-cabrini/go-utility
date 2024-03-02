// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"fmt"
	"runtime"
)

// Mypanic
// If err is not nil logs a FATAL error message
func Mypanic(err error) {
	if err == nil {
		return
	}

	Logf(FATAL, "%v", err)
}

/*
Takes a funcion `f` that might return an error, a void function `pre`, a void function `post`

It is supposed to be used to handle deferred function calls that returns an error without ugly sintax such as:

	defer func() {
		err := fp.Close()

		if err != nil {
			panic(err)
		}
	}()

Using Deferrable:

	defer Deferrable(fp.Close, nil, nil)

What it does:

  - If pre is not nil, executes pre();
  - Executes f() and panic (using Mypanic) if `f` result is not nil;
  - If does not panic and post is not nil, executes post.
*/
func Deferrable(f func() error, pre func(), post func()) {
	if pre != nil {
		pre()
	}
	Mypanic(f())
	if post != nil {
		post()
	}
}

// Appends to error the funcion in wich the error has been generated
func AppendError(err error) error {
	if err == nil {
		return nil
	}
	var fileName = trace()[1]
	return fmt.Errorf("%s: %v", fileName, err)
}

// Returns the third to last function on the stack
// Should be used only in AppendError
func trace() (receipts []string) {
	var p = make([]uintptr, 10)
	runtime.Callers(2, p)

	for _, pcx := range p {
		f := runtime.FuncForPC(pcx)
		receipts = append(receipts, f.Name())
	}

	return
}
