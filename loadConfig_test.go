package loadConfig_test

import (
	"fmt"
	"testing"

	"github.com/dan0505/loadConfig"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	v := viper.New()
	v.AddConfigPath(".")

	v.SetConfigName("test")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error while reading config file %s", err))
	}

	t.Run("panic if cfg is not set", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("The code did not panic")
			}
			assert.Equal(t, r, "tag not set for TestString")
		}()
		type panicStruct struct {
			TestString string
		}
		var testStruct panicStruct
		loadConfig.LoadConfig(&testStruct, v)
		t.Error("should panic")
	})

	t.Run("ignore if set -", func(t *testing.T) {
		type stringStruct struct {
			TestString string `cfg:"-"`
		}
		var testStruct stringStruct
		loadConfig.LoadConfig(&testStruct, v)
		assert.Equal(t, testStruct.TestString, "")
	})

	t.Run("load string", func(t *testing.T) {
		type stringStruct struct {
			TestString string `cfg:"load_string"`
		}
		var testStruct stringStruct
		loadConfig.LoadConfig(&testStruct, v)
		assert.Equal(t, testStruct.TestString, "test string")
	})

	t.Run("panic if no value", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("The code did not panic")
			}
			assert.Equal(t, r, "can't get config for fake_entry")
		}()
		type panicStruct struct {
			TestString string `cfg:"fake_entry"`
		}
		var testStruct panicStruct
		loadConfig.LoadConfig(&testStruct, v)
		t.Error("should panic")
	})

	t.Run("load int", func(t *testing.T) {
		type intStruct struct {
			TestString int `cfg:"load_int"`
		}
		var testStruct intStruct
		loadConfig.LoadConfig(&testStruct, v)
		assert.Equal(t, testStruct.TestString, 100)
	})

	t.Run("load struct", func(t *testing.T) {
		type structStruct struct {
			TestStruct struct {
				TestString string `cfg:"load_string"`
				TestInt    int    `cfg:"load_int"`
			} `cfg:"struct"`
		}
		var testStruct structStruct
		loadConfig.LoadConfig(&testStruct, v)
		assert.Equal(t, "test string 2", testStruct.TestStruct.TestString)
		assert.Equal(t, 200, testStruct.TestStruct.TestInt)
	})

	t.Run("should ignore - struct", func(t *testing.T) {
		type RandomInterface interface{}
		type RandomStruct struct{}
		type structStruct struct {
			RandomPtr     *RandomInterface `cfg:"-"`
			AnotherPtr    *structStruct    `cfg:"-"`
			AnotherStruct RandomStruct     `cfg:"-"`
		}
		var testStruct structStruct
		loadConfig.LoadConfig(&testStruct, v)
	})
}

func TestLoadConfigWithEnvEntry(t *testing.T) {
	v := viper.New()
	v.AddConfigPath(".")

	v.SetConfigName("test")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error while reading config file %s", err))
	}

	t.Run("load with entry name", func(t *testing.T) {
		type structStruct struct {
			TestString string `cfg:"load_string"`
			TestInt    int    `cfg:"load_int"`
		}
		var testStruct structStruct
		loadConfig.LoadConfigWithEnvEntry(&testStruct, v, "struct")
		assert.EqualValues(t, structStruct{"test string 2", 200}, testStruct)
	})
}
