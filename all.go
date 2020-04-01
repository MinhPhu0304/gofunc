package gofunc

import (
	"reflect"
)

// All return true if all elements of the list match the predicate
func All(value interface{}) func(list interface{}) bool {

	return func(list interface{}) bool {
		switch list.(type) {
		case []int, []string, []float32, []float64, []bool, []complex64, []complex128, []int64, []uint64:
			valueType := reflect.ValueOf(value).Interface()
			for index := 0; index < reflect.ValueOf(list).Len(); index++ {
				if valueType != reflect.ValueOf(list).Index(index).Interface() {
					return false
				}
			}
			return true
		default:
			return false
		}
	}
}
