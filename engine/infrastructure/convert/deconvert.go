package convert

import (
	"reflect"
	"strings"
)

func DeConvert[TFrom any, TTo any](data TFrom) TTo {
	v := reflect.ValueOf(data)

	var result TTo
	fields := make(map[string]any, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		if !isValid(v, i) {
			continue
		}

		tag := getTag(v, i, "convert")
		var fieldName string

		if tag != "" {
			if checkInTag(tag, "func") {
				continue
			}

			values := strings.Split(tag, ",")
			fieldName = values[0]
		} else {
			fieldName = reflect.Indirect(v).Type().Field(i).Name
		}

		fieldValue := v.Field(i).Interface()
		fields[fieldName] = fieldValue
	}

	vR := reflect.ValueOf(result)

	for i := 0; i < vR.NumField(); i++ {
		if !isValid(vR, i) {
			continue
		}

		fieldName := reflect.Indirect(vR).Type().Field(i).Name

		if _, ok := fields[fieldName]; !ok {
			continue
		}

		value := reflect.ValueOf(fields[fieldName])
		reflect.ValueOf(&result).Elem().Field(i).Set(value)
	}

	return result
}

func DeConvertList[TFrom any, TTo any](data []TFrom) []TTo {
	var result []TTo

	for _, item := range data {
		result = append(result, DeConvert[TFrom, TTo](item))
	}

	return result
}
