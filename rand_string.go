// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"crypto/rand"
	"fmt"
)

func RandString(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("lunghezza non valida: %d", n)
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const mask = 63
	const maxByte = 255 - (255 % len(charset))

	b := make([]byte, n)
	i := 0
	for i < n {
		buf := make([]byte, 1)
		if _, err := rand.Read(buf); err != nil {
			return "", AppendError(err)
		}

		if int(buf[0]) > maxByte {
			continue
		}

		idx := int(buf[0] & mask)
		if idx < len(charset) {
			b[i] = charset[idx]
			i++
		}
	}

	return string(b), nil
}
