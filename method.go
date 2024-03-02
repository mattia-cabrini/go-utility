// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"reflect"
)

// Represent an object method, callable by invoking F and passing the very same
// parameters that would be passed to that object method (which means NOT the receiver)
type Method struct {
	to    reflect.Method
	numIn int

	F GenericFunc
}

func newMethod(to reflect.Method, callable GenericFunc) (m *Method) {
	m = &Method{
		to:    to,
		numIn: to.Type.NumIn() - 1,
		F:     callable,
	}

	return
}

// Retuns the numer of parameters that F expects
func (m *Method) NumIn() int {
	return m.numIn
}

/*
Takes an integer `i` and retuns the i-th parameter kind that F would expect.

Note that F do take as parameters the object method's parameters but NOT the method receiver
*/
func (m *Method) ParamKind(i int) reflect.Kind {
	return m.to.Type.In(i + 1).Kind()
}
