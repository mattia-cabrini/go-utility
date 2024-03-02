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

/*
Lock a lockable and retun a funcion to unlock the lockable.

The meaning is to avoid ugly syntax such as:

	mu.Lock()
	defer mu.Unlock()

Which became:

	defer Monitor(mu)()
*/
func Monitor(mu lockable) func() {
	mu.Lock()
	return func() {
		mu.Unlock()
	}
}

/*
Lock a rlockable and retun a funcion to unlock the rlockable.

The meaning is to avoid ugly syntax such as:

	mu.RLock()
	defer mu.RUnlock()

Which became:

	defer RMonitor(mu)()
*/
func RMonitor(mu rlockable) func() {
	mu.RLock()
	return func() {
		mu.RUnlock()
	}
}
