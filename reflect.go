// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"fmt"
	"reflect"
)

type GenericFunc func(...interface{}) ([]interface{}, error)

func checkMethodArgs(methodTO reflect.Method, args []reflect.Value) error {
	if expected, actual := methodTO.Type.NumIn()-1, len(args); expected != actual {
		return fmt.Errorf("method %s expected %d arguments, got %d", methodTO.Name, expected, actual)
	}

	for i, arg := range args {
		argTypeExpected := methodTO.Type.In(i + 1)

		if !arg.Type().AssignableTo(argTypeExpected) {
			return fmt.Errorf("method %s expected %s argument, got %s (parameter #%d)", methodTO.Name, argTypeExpected.String(), arg.String(), i+1)
		}
	}

	return nil
}

func newGenericFunc(methodTO reflect.Method, methodVO reflect.Value) GenericFunc {
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

// Giveng an obj interface{}, returns a *Method from that obj.
// GetMethod seeks a method that has name equal to name + suffix.
// `suffix` might just be an empty string.
func GetMethod(obj interface{}, name string, suffix string) *Method {
	to := reflect.TypeOf(obj)
	vo := reflect.ValueOf(obj)

	methodTO, b := to.MethodByName(name + suffix)

	if b {
		methodVO := vo.MethodByName(name + suffix)
		callable := newGenericFunc(methodTO, methodVO)
		return newMethod(methodTO, callable)
	}

	return nil
}

// Giveng an obj interface{}, returns an interface{} which is a propery value of that obj.
// GetProperty seeks a property that has name equal to name + suffix and all given tags equal to true.
// `suffix` might just be an empty string.
func GetProperty(obj interface{}, name string, suffix string, tags ...string) interface{} {
	to := reflect.TypeOf(obj)
	vo := reflect.ValueOf(obj)

	fieldTO, b := to.FieldByName(name + suffix)

	if b {
		for _, tx := range tags {
			if tagValue := fieldTO.Tag.Get(tx); tagValue != "true" {
				b = false
				break
			}
		}

		if b {
			fieldVO := vo.FieldByName(name + suffix)

			return fieldVO.Interface()
		}
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
