package loadConfig

import (
	"fmt"
	"reflect"
	"strings"
)

func isPtrToStruct(f interface{}) error {
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

func combineTage(lastTag, newTag string) string {
	tags := []string{}
	tags = append(tags, strings.Split(lastTag, ".")...)
	tags = append(tags, newTag)
	nonEmptyTags := []string{}
	for _, t := range tags {
		if t != "" {
			nonEmptyTags = append(nonEmptyTags, t)
		}
	}
	return strings.Join(nonEmptyTags, ".")
}
