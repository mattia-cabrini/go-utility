// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"fmt"
	"os"
	"os/exec"
)

func Lessf(format string, args ...interface{}) error {
	str := fmt.Sprintf(format, args...)
	return less(str)
}

func less(str string) (err error) {
	echo := exec.Command("echo", str)
	echo.Stdin = os.Stdin
	echo.Stderr = os.Stderr

	lessCmd := exec.Command("less")
	lessCmd.Stdout = os.Stdout

	if lessCmd.Stdin, err = echo.StdoutPipe(); err != nil {
		return
	}

	if err = lessCmd.Start(); err != nil {
		return
	}
	if err = echo.Run(); err != nil {
		return
	}
	if err = lessCmd.Wait(); err != nil {
		return
	}

	return
}
