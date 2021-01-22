package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type fun func()

func (f *fun) Print(s string) {
	fmt.Println(s)
}

func TestFun(t *testing.T) {
	var f fun
	f.Print("1")
	typ := reflect.TypeOf(&f)
	//typ := reflect.TypeOf(f)
	fmt.Println(typ.Name(), typ.Kind())
	method, ok := typ.MethodByName("Print")
	require.True(t, ok)
	fmt.Println(method)

	f1 := reflect.TypeOf(nil)
	fmt.Println(f1)
}

func TestBatchCall(t *testing.T) {
	err := BatchCallByReflect(2, foo1, "", []int{1, 2})
	require.NoError(t, err)
	err = BatchCallByReflect(2, foo1, "")
	require.Error(t, err)
	err = BatchCallByReflect(2, foo2, "", []int{1, 2})
	require.Equal(t, err, Err)
}

// BenchmarkBatch-12    	 2363384	       464 ns/op
func BenchmarkBatchCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BatchCallByReflect(2, foo1, "", []int{1, 2, 3, 4})
	}
}

// BenchmarkCall-12    	201063060	         5.37 ns/op
func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BatchFoo1(2, "", []int{1, 2, 3, 4})
	}
}
