// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"os"
	"os/exec"
)

func Editor(editorPath string, path string) (err error) {
	editor := exec.Command(editorPath, path)
	editor.Stdin = os.Stdin
	editor.Stderr = os.Stderr
	editor.Stdout = os.Stdout

	if err = editor.Run(); err != nil {
		return
	}

	return
}
