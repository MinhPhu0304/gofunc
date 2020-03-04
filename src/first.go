package gofunc

import "reflect"

// First : Get the first element in the array
func First(input interface{}) interface{} {
	if reflect.ValueOf(input).Kind() != reflect.Array {
		return nil
	}
	result := reflect.ValueOf(input).Index(0)
	return result
}
