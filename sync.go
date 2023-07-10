// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

type lockable interface {
	Lock()
	Unlock()
}

type rlockable interface {
	RLock()
	RUnlock()
}

func Monitor(mu lockable) func() {
	mu.Lock()
	return func() {
		mu.Unlock()
	}
}

func RMonitor(mu rlockable) func() {
	mu.RLock()
	return func() {
		mu.RUnlock()
	}
}
