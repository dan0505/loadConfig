package loadConfig

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

func LoadConfig(c interface{}, v *viper.Viper) {
	if err := IsPtrToStruct(c); err != nil {
		panic(err)
	}
	container := reflect.ValueOf(c).Elem()
	setValue(container, v, "")

}

func setValue(container reflect.Value, v *viper.Viper, tag string) {
	for j := 0; j < container.NumField(); j++ {
		fieldType := container.Type().Field(j)
		fieldName := fieldType.Name
		tag := string(fieldType.Tag.Get("cfg"))
		if tag == "" {
			panic(fmt.Sprintf("tag not set for %s", fieldName))
		}
		if tag == "-" {
			continue
		}

		f := container.Field(j)
		if !f.IsValid() || !f.CanSet() {
			panic(fmt.Sprintf("can't set field %s", fieldName))
		}

		noValue := func() {
			panic(fmt.Sprintf("can't get config for %s", tag))
		}
		switch f.Type().Kind() {
		case reflect.String:
			value := v.GetString(tag)
			if value == "" {
				noValue()
			}
			f.SetString(value)
		case reflect.Int:
			value := int64(v.GetInt(tag))
			if value == 0 {
				noValue()
			}
			f.SetInt(int64(v.GetInt(tag)))
		case reflect.Struct:
			setValue(f, v, tag)
		default:
			panic(fmt.Sprintf("can't handle %s", f.Type().String()))
		}
	}
}
