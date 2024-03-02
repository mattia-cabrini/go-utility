// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"os"
	"os/exec"
)

/*
To run an editor on the terminal.

Editor merely runs the editor program (param editorPath) by passing as first parameter a file path (param path)

Parameters:

  - editorPath string: the path to the editor binary;
  - path string: the path to wich the edited file will be saved.

Returns:

  - err error: an error if exec.Command.Run shoud generate one.
*/
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
