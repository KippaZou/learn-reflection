package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string
}

func (c *Person) String() string {
	return c.Name
}

func TestReflectionLaw1(t *testing.T) {
	coder := &Person{Name: "kippa"}
	typ := reflect.TypeOf(coder)
	val := reflect.ValueOf(coder)
	typeOfStringer := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println("kind of coder:", typ.Kind())
	fmt.Println("TypeOf coder:", typ)
	fmt.Println("ValueOf coder:", val)
	//fmt.Println("field 0 of coder:", typ.Field(0))
	fmt.Println("implements stringer:", typ.Implements(typeOfStringer))
}

func TestReflectionLaw2(t *testing.T) {
	coder := &Person{Name: "kippa"}
	val := reflect.ValueOf(coder)
	c, ok := val.Interface().(*Person)
	if ok {
		fmt.Println(c.Name)
	} else {
		panic("type assert to *Person err")
	}
}

// panic
func TestReflectionLaw3V1(t *testing.T) {
	i := 1
	v := reflect.ValueOf(i)
	v.SetInt(10)
	fmt.Println(i)
}

/* 通过获取变量的指针来更新值
相当于
	i := 1
	v := &i
	*v = 10
*/
func TestReflectionLaw3V2(t *testing.T) {
	i := 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)
}
