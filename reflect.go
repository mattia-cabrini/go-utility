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

import (
	"fmt"
	"reflect"
)

type genericFunc func(...interface{}) ([]interface{}, error)

func checkMethodArgs(methodVO reflect.Method, args []reflect.Value) error {
	if m, a := methodVO.Type.NumIn(), len(args); m != a+1 {
		return fmt.Errorf("expected %d arguments, got %d", m, a)
	}

	for i, arg := range args {
		argTypeExpected := methodVO.Type.In(i + 1)

		if i == 0 {
			continue // skip the object itself
		}

		if !arg.Type().AssignableTo(argTypeExpected) {
			return fmt.Errorf("expected %v argument, got %v (parameter #%d)", argTypeExpected.Kind(), arg.Kind(), i)
		}
	}

	return nil
}

func newGenericFunc(obj interface{}, methodTO reflect.Method, methodVO reflect.Value) genericFunc {
	return func(argsI ...interface{}) (returnValues []interface{}, argError error) {
		var args = make([]reflect.Value, len(argsI))
		for i, arg := range argsI {
			args[i] = reflect.ValueOf(arg)
		}

		if argError = checkMethodArgs(methodTO, args); argError != nil {
			return
		}

		returnValuesVO := methodVO.Call(args)

		returnValues = make([]interface{}, len(returnValuesVO))
		for i, returnValueVO := range returnValuesVO {
			returnValues[i] = returnValueVO.Interface()
		}

		return
	}
}

func GetMethod(obj interface{}, name string, suffix string) genericFunc {
	to := reflect.TypeOf(obj)
	vo := reflect.ValueOf(obj)

	methodTO, b := to.MethodByName(name + suffix)

	if b {
		methodVO := vo.MethodByName(name + suffix)

		return newGenericFunc(obj, methodTO, methodVO)
	}

	return nil
}

var typeDefault = map[reflect.Kind]interface{}{
	reflect.String:  "",
	reflect.Int64:   int64(0),
	reflect.Int32:   int32(0),
	reflect.Int16:   int16(0),
	reflect.Int:     int(0),
	reflect.Float64: float64(0),
	reflect.Float32: float32(0),
}

func TypeDefault(k reflect.Kind) (v interface{}, b bool) {
	v, b = typeDefault[k]
	return
}

func IPtrToI(fv reflect.Value, i interface{}) (a interface{}) {
	var diPtr = i.(*interface{})
	var di = *diPtr

	switch fv.Kind() {
	case reflect.String:
		a = di.(string)
	case reflect.Int64:
		a = di.(*int64)
	case reflect.Int16:
		a = di.(*int16)
	case reflect.Int32:
		a = di.(*int32)
	case reflect.Int:
		a = di.(*int)
	case reflect.Float32:
		a = di.(*float32)
	case reflect.Float64:
		a = di.(*float64)
	}

	return
}
