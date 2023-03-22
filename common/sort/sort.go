package sort

import (
	"reflect"
	"sort"
	"strings"

	"github.com/yearn/ydaemon/common/logs"
)

func SortBy[T any](jsonField string, sortOrder string, arr []T) {
	if len(arr) == 0 {
		return
	}
	valueType := reflect.TypeOf(arr[0])
	jsonFields := strings.Split(jsonField, ".")
	jsonFieldsLen := len(jsonFields)
	fields := make([]reflect.StructField, jsonFieldsLen)
	match := 0

	for _, jsonField := range jsonFields {
		if valueType.Kind() == reflect.Struct {
			for i := 0; i <= valueType.NumField(); i++ {
				if (i == valueType.NumField()) && (match == 0) {
					return
				}
				fields[match] = valueType.Field(i)

				if fields[match].Tag.Get("json") == jsonField || strings.EqualFold(fields[match].Name, jsonField) {
					valueType = fields[match].Type
					if valueType.Kind() == reflect.Pointer {
						valueType = valueType.Elem()
					}
					match++
					if match == jsonFieldsLen {
						break
					}
				}
			}
		}
		if match == jsonFieldsLen {
			break
		}
	}
	if match == 0 || match != jsonFieldsLen {
		logs.Error("field not found")
		return
	}

	sort.Slice(arr, func(i, j int) bool {
		lastField := reflect.StructField{}
		v1 := reflect.ValueOf(arr[i])
		v2 := reflect.ValueOf(arr[j])
		for _, field := range fields {
			if v1.Kind() == reflect.Pointer {
				v1 = v1.Elem()
			}
			if v2.Kind() == reflect.Pointer {
				v2 = v2.Elem()
			}
			v1 = v1.FieldByName(field.Name)
			v2 = v2.FieldByName(field.Name)
			lastField = field
		}

		switch lastField.Type.Name() {
		case "int", "int8", "int16", "int32", "int64":
			if sortOrder == "asc" {
				return v1.Int() < v2.Int()
			} else {
				return v1.Int() > v2.Int()
			}
		case "uint", "uint8", "uint16", "uint32", "uint64":
			if sortOrder == "asc" {
				return v1.Uint() < v2.Uint()
			} else {
				return v1.Uint() > v2.Uint()
			}
		case "float32", "float64":
			if sortOrder == "asc" {
				return v1.Float() < v2.Float()
			} else {
				return v1.Float() > v2.Float()
			}
		case "string", "MixedcaseAddress", "Address":
			if sortOrder == "asc" {
				return v1.String() < v2.String()
			} else {
				return v1.String() > v2.String()
			}
		case "bool":
			if sortOrder == "asc" {
				return !v1.Bool() // return small numbers first
			} else {
				return v1.Bool() // return big numbers first
			}
		default:
			logs.Warning(`field type [` + lastField.Type.Name() + `] not supported`)
			return false
		}
	})
}
