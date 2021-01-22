package test

import (
	"fmt"
	"reflect"
)

func Sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("types of elements are not the same")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

func MakeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), Sum)

	fn.Set(v)
}
