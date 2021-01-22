package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceEqual(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	//fmt.Println(slice1 == slice2)
	fmt.Println(reflect.DeepEqual(slice1, slice2))
}

func TestStructEqual(t *testing.T) {
	p1 := &Person{Name: "kippa"}
	p2 := &Person{Name: "kippa"}
	fmt.Println(p1 == p2)
	fmt.Println(reflect.DeepEqual(p1, p2))
	// 不用反射可以这样
	fmt.Println(*p1 == *p2)
	type Coder struct {
		*Person
		Language []string
	}
	c1 := &Coder{
		Person:   p1,
		Language: []string{"go", "python", "c"},
	}
	c2 := &Coder{
		Person:   p1,
		Language: []string{"go", "python", "c"},
	}
	//fmt.Println(*c1 == *c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}
