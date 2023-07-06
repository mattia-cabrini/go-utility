/*
MIT License

Copyright (c) 2023 Mattia Cabrini

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package utility

import "testing"

type testType struct {
	A int
	B int
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

	f := GetMethod(&obj, "GetAB", "")

	if f == nil {
		t.Fail()
		return
	}

	results, err := f(0, 1)

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

func TestMethodOk1(t *testing.T) {
	obj := testType{0, 1}

	f := GetMethod(&obj, "GetAB", "interface")

	if f == nil {
		t.Fail()
		return
	}

	results, err := f(0, 1)

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

	f := GetMethod(&obj, "GetAB", "struct")

	if f == nil {
		t.Fail()
		return
	}

	results, err := f(0, obj)

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

	f := GetMethod(&obj, "Get", "AB")

	if f == nil {
		t.Fail()
		return
	}

	_, err := f(0)

	if err == nil {
		t.Fail()
	}
}

func TestMethodArgTypeKo(t *testing.T) {
	obj := testType{0, 1}

	f := GetMethod(&obj, "Get", "AB")

	if f == nil {
		t.Fail()
		return
	}

	_, err := f(0, "1")

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
