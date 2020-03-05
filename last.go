package gofunc

import "reflect"

// Last : Get the Last element in the array
func Last(input interface{}) interface{} {
	if reflect.ValueOf(input).Kind() != reflect.Array {
		return nil
	}
	result := reflect.ValueOf(input)

	return result.Index(result.Len() - 1).Interface()
}
