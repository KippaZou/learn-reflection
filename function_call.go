package test

import (
	"fmt"
	"reflect"
)

var Err = fmt.Errorf("an error")

func foo1(_ string, _ []int)       {}
func foo2(_ string, _ []int) error { return Err }

func BatchFoo2(limit int, str string, ints []int) error {
	for index := 0; index*limit < len(ints); index++ {
		if (index+1)*limit < len(ints) {
			err := foo2(str, ints[index*limit:(index+1)*limit])
			if err != nil {
				return err
			}
		} else {
			err := foo2(str, ints[index*limit:])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func BatchFoo1(limit int, str string, ints []int) {
	for index := 0; index*limit < len(ints); index++ {
		if (index+1)*limit < len(ints) {
			foo1(str, ints[index*limit:(index+1)*limit])
		} else {
			foo1(str, ints[index*limit:])
		}
	}
}

// BatchCallByReflect
func BatchCallByReflect(
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
	call := func(argvIn []reflect.Value) error {
		retValues := function.Call(argvIn)
		if len(retValues) == 0 {
			return nil
		}
		for i := len(retValues) - 1; i >= 0; i-- {
			if err, ok := retValues[i].Interface().(error); ok {
				return err
			}
		}
		return nil
	}

	for index := 0; index*limit < sliceValueLen; index++ {
		if (index+1)*limit < sliceValueLen {
			argvIn[argsLen-1] = sliceValue.Slice(index*limit, (index+1)*limit)
		} else {
			argvIn[argsLen-1] = sliceValue.Slice(index*limit, sliceValueLen)
		}
		err := call(argvIn)
		if err != nil {
			return err
		}
	}
	return nil
}
