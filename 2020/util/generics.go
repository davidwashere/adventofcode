package util

import "reflect"

// convToType will return a 'reflected' value of the kind specified
// ref: https://ahmet.im/blog/golang-take-slices-of-any-type-as-input-parameter/
func convToType(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
