package test

import (
	"fmt"
	"reflect"
)

func BatchCall(
	limit int,
	f interface{},
	args ...interface{},
) error {
	argsLen := len(args)
	if argsLen == 0 {
		return fmt.Errorf("args len cant be zero")
	}
	if limit < 1 {
		return fmt.Errorf("limit < 1")
	}
	sliceValue := reflect.ValueOf(args[argsLen-1])
	if sliceValue.Kind() != reflect.Slice {
		return fmt.Errorf("last arg must be slice")
	}
	function := reflect.ValueOf(f)
	if function.Kind() != reflect.Func {
		return fmt.Errorf("invalid func kind:%s", sliceValue.Kind().String())
	}
	if argsLen != function.Type().NumIn() {
		return fmt.Errorf("wrong args len, expected: %d, got: %d", function.Type().NumIn(), argsLen)
	}
	argv := make([]reflect.Value, argsLen)
	for i := 0; i < argsLen-1; i++ {
		argv[i] = reflect.ValueOf(args[i])
	}
	argvIn := append(argv[:argsLen-1], reflect.Value{})
	sliceValueLen := sliceValue.Len()
	for index := 0; index*limit < sliceValueLen; index++ {
		if (index+1)*limit < sliceValueLen {
			argvIn[argsLen-1] = sliceValue.Slice(index*limit, (index+1)*limit)
		} else {
			argvIn[argsLen-1] = sliceValue.Slice(index*limit, sliceValueLen)
		}
		function.Call(argvIn)
	}
	return nil
}
