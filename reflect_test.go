// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"reflect"
	"testing"
)

type testType struct {
	A int `con:"true" con2:"true"`
	B int `con:"true" con2:"false"`
}

func (t *testType) GetAB(a, b int) (int, int) {
	return a + t.A, b + t.B
}

func (t *testType) GetABinterface(a int, b interface{}) (int, int) {
	return a + t.A, 0
}

func (t *testType) GetABstruct(a int, b testType) (int, int) {
	return a + t.A, t.B + b.B
}

func TestMethodOk0(t *testing.T) {
	obj := testType{0, 1}

	m := GetMethod(&obj, "GetAB", "")

	if m == nil {
		t.Fail()
		return
	}

	results, err := m.F(0, 1)

	if err != nil {
		t.Error(err)
		return
	}

	if results[0].(int) != 0 {
		t.Error(0)
	}
	if results[1].(int) != 2 {
		t.Error(1)
	}

	if m.ParamKind(0) != reflect.Int {
		t.Error(2)
	}
	if m.ParamKind(1) != reflect.Int {
		t.Error(3)
	}
	if m.NumIn() != 2 {
		t.Error(4)
	}
}

func TestMethodOk1(t *testing.T) {
	obj := testType{0, 1}

	m := GetMethod(&obj, "GetAB", "interface")

	if m == nil {
		t.Fail()
		return
	}

	results, err := m.F(0, 1)

	if err != nil {
		t.Error(err)
		return
	}

	if results[0].(int) != 0 {
		t.Error(0)
	}
	if results[1].(int) != 0 {
		t.Error(1)
	}
}

func TestMethodOk2(t *testing.T) {
	obj := testType{0, 1}

	m := GetMethod(&obj, "GetAB", "struct")

	if m == nil {
		t.Fail()
		return
	}

	results, err := m.F(0, obj)

	if err != nil {
		t.Error(err)
		return
	}

	if results[0].(int) != 0 {
		t.Error(0)
	}
	if results[1].(int) != 2 {
		t.Error(1)
	}
}

func TestMethodArgCountKo(t *testing.T) {
	obj := testType{0, 1}

	m := GetMethod(&obj, "Get", "AB")

	if m == nil {
		t.Fail()
		return
	}

	_, err := m.F(0)

	if err == nil {
		t.Fail()
	}
}

func TestMethodArgTypeKo(t *testing.T) {
	obj := testType{0, 1}

	m := GetMethod(&obj, "Get", "AB")

	if m == nil {
		t.Fail()
		return
	}

	_, err := m.F(0, "1")

	if err == nil {
		t.Fail()
	}
}

func TestMethodNameKo(t *testing.T) {
	obj := testType{0, 1}

	f := GetMethod(&obj, "Get2", "AB")

	if f != nil {
		t.Fail()
	}
}

func TestFieldOk(t *testing.T) {
	obj := testType{0, 1}

	i := GetProperty(obj, "B", "")

	if i.(int) != 1 {
		t.Fail()
	}
}

func TestFieldTag(t *testing.T) {
	obj := testType{0, 1}

	i := GetProperty(obj, "B", "", "con", "con2")

	if i != nil {
		t.Fail()
	}

	i = GetProperty(obj, "A", "", "con", "con2")

	if i.(int) != 0 {
		t.Fail()
	}
}

func TestFieldKo(t *testing.T) {
	obj := testType{0, 1}

	i := GetProperty(obj, "c", "")

	if i != nil {
		t.Fail()
	}
}
