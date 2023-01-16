package loadConfig

import (
	"fmt"
	"reflect"
)

func IsPtrToStruct(f interface{}) error {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("interface %s is not a pointer", v.Type().Name())
	}
	v = v.Elem() // element behind the ptr
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("interface in pointer %s is not a struct", v.Type().Name())
	}
	return nil
}
