package field

import (
	"reflect"
)

func IterPublicFields(record interface{}) []WrappedField {
	reflected := getReflectedRecord(record)
	fields := []WrappedField{}
	for i := 0; i < reflected.NumField(); i++ {
		field := GetFieldWithIndex(reflected, i)
		if field.IsPublic() {
			fields = append(fields, field)
		}
	}
	return fields
}

func getReflectedRecord(record interface{}) reflect.Value {
	switch recordType := record.(type) {
	case reflect.Value:
		return recordType
	default:
		reflected := reflect.ValueOf(record)
		switch reflected.Kind() {
		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			return reflected.Elem()
		default:
			return reflected
		}
	}
}

func GetFieldWithIndex(record interface{}, i int) WrappedField {
	reflected := getReflectedRecord(record)
	return WrappedField{
		StructField: reflected.Type().Field(i),
		Value:       reflected.Field(i),
	}
}

func GetFieldWithName(record interface{}, name string) (wrappedField WrappedField, ok bool) {
	reflected := getReflectedRecord(record)
	var structField reflect.StructField
	if structField, ok = reflected.Type().FieldByName(name); ok {
		wrappedField = WrappedField{
			StructField: structField,
			Value:       reflected.FieldByName(name),
		}
	}
	return
}
