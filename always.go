package gofunc

import "reflect"

// Always : A curried function that returns functio which always return value supplied at the start.
func Always(value interface{}) func() interface{} {
	return func() interface{} {
		return reflect.ValueOf(value).Interface()
	}
}
