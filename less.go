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
