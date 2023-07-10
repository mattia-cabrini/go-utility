// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"reflect"
)

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

func (m *Method) NumIn() int {
	return m.numIn
}

func (m *Method) ParamKind(i int) reflect.Kind {
	return m.to.Type.In(i + 1).Kind()
}
