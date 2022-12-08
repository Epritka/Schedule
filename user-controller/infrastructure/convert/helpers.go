package convert

import (
	"reflect"
	"strings"
)

func checkInTag(tag string, name string) bool {
	values := strings.Split(tag, ",")
	for i, v := range values {
		if i == 0 {
			continue
		}

		if v == name {
			return true
		}
	}
	return false
}

func getTag(v reflect.Value, index int, name string) string {
	return reflect.Indirect(v).Type().Field(index).Tag.Get("convert")
}

func isValid(v reflect.Value, index int) bool {
	// Проверяем что свойстов публичное
	if !reflect.Indirect(v).Type().Field(index).IsExported() {
		return false
	}

	// проверяем что это не вложенная структура или slice
	kind := v.Field(index).Kind()
	if kind == reflect.Struct || kind == reflect.Slice {
		return false
	}

	return true
}
