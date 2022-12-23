package convert

import (
	"reflect"
	"strings"
)

func Convert[TFrom any, TTo any](data TFrom) TTo {
	v := reflect.ValueOf(data)

	var result TTo
	fields := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		if !isValid(v, i) {
			continue
		}

		fieldName := reflect.Indirect(v).Type().Field(i).Name
		fieldValue := v.Field(i).Interface()
		fields[fieldName] = fieldValue
	}

	vR := reflect.ValueOf(result)

	for i := 0; i < vR.NumField(); i++ {
		if !isValid(vR, i) {
			continue
		}

		tag := getTag(vR, i, "convert")
		var fieldName string
		var value reflect.Value

		if tag != "" {
			values := strings.Split(tag, ",")
			fieldName = values[0]

			if checkInTag(tag, "func") {
				params := []reflect.Value{
					reflect.ValueOf(data),
				}
				value = reflect.ValueOf(&result).MethodByName(fieldName).Call(params)[0]
			} else {
				value = reflect.ValueOf(fields[fieldName])
			}

		} else {
			fieldName = reflect.Indirect(vR).Type().Field(i).Name
			value = reflect.ValueOf(fields[fieldName])
		}

		reflect.ValueOf(&result).Elem().Field(i).Set(value)
	}

	return result
}

func ConvertList[TFrom any, TTo any](data []TFrom) []TTo {
	var result []TTo

	for _, item := range data {
		result = append(result, Convert[TFrom, TTo](item))
	}

	return result
}
