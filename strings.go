// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import "strings"

// Returns str without last n characters
func ButLastN(str string, n int) string {
	lenStr := len(str)

	if lenStr < n {
		return ""
	}

	return str[:lenStr-n]
}

// Returns true is str ends with any of the provided suffixes; False otherwise
// If true is returned, the matching suffix is also returned; Empty string otherwise
func EndsWithAny(str string, suffixesList ...string) (bool, string) {
	for _, suffix := range suffixesList {
		if EndsWith(str, suffix) {
			return true, suffix
		}
	}

	return false, ""
}

// Returns true is str ends with suffix; False otherwise
func EndsWith(str string, suffix string) bool {
	lenStr := len(str)
	lenSuffix := len(suffix)

	if lenStr < lenSuffix {
		return false
	}

	return str[lenStr-lenSuffix:] == suffix
}

func Quote(str string) string {
	var sb strings.Builder

	sb.WriteRune('"')

	for _, rx := range str {
		if rx == '"' {
			sb.WriteString("\\\"")
		} else {
			sb.WriteRune(rx)
		}
	}

	sb.WriteRune('"')

	return sb.String()
}
