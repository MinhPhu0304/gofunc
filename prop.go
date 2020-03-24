package gofunc

import (
	"reflect"
	"strings"
)

// Prop : Return a field with given name from a struct. Note: all field should be export otherwise the function won't work
func Prop(propertiesName string, obj interface{}) interface{} {
	objType := reflect.TypeOf(obj)

	if strings.Contains(objType.String(), ".") {
		return reflect.ValueOf(obj).FieldByName(propertiesName).Interface()
	}
	return nil
}
