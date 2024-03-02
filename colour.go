// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

var (
	Reset  string = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Purple        = "\033[35m"
	Cyan          = "\033[36m"
	Gray          = "\033[37m"
	White         = "\033[97m"
)

// Set colours strings to ""
func NotInteractive() {
	Reset = ""
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Purple = ""
	Cyan = ""
	Gray = ""
	White = ""
}

// Reset colours strings to their values
func Interactive() {
	Reset = "\033[0m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Purple = "\033[35m"
	Cyan = "\033[36m"
	Gray = "\033[37m"
	White = "\033[97m"
}
