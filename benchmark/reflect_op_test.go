// referred from https://colobu.com/2019/01/29/go-reflect-performance/
package benchmark

import (
	"reflect"
	"testing"
)

type Student struct {
	Name  string
	Age   int
	Class string
	Score int
}

// BenchmarkReflect_New-12                 27316540                43.8 ns/op
func BenchmarkReflect_New(b *testing.B) {
	var s *Student
	sv := reflect.TypeOf(Student{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv)
		s, _ = sn.Interface().(*Student)
	}
	_ = s
}

// BenchmarkDirect_New-12                  39993478                29.7 ns/op
func BenchmarkDirect_New(b *testing.B) {
	var s *Student
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = new(Student)
	}
	_ = s
}

// BenchmarkReflect_SetFieldByName-12       3993316               306 ns/op
func BenchmarkReflect_SetFieldByName(b *testing.B) {
	sv := reflect.TypeOf(Student{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv).Elem()
		sn.FieldByName("Name").SetString("Jerry")
		sn.FieldByName("Age").SetInt(18)
		sn.FieldByName("Class").SetString("20005")
		sn.FieldByName("Score").SetInt(100)
	}
}

// BenchmarkReflect_SetFieldByIndex-12     17714866                66.6 ns/op
func BenchmarkReflect_SetFieldByIndex(b *testing.B) {
	sv := reflect.TypeOf(Student{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv).Elem()
		sn.Field(0).SetString("Jerry")
		sn.Field(1).SetInt(18)
		sn.Field(2).SetString("20005")
		sn.Field(3).SetInt(100)
	}
}

// BenchmarkReflect_Set-12                 26631484                45.5 ns/op
func BenchmarkReflect_Set(b *testing.B) {
	var s *Student
	sv := reflect.TypeOf(Student{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv)
		s = sn.Interface().(*Student)
		s.Name = "Jerry"
		s.Age = 18
		s.Class = "20005"
		s.Score = 100
	}
}

// BenchmarkDirect_Set-12                  38555898                30.8 ns/op
func BenchmarkDirect_Set(b *testing.B) {
	var s *Student
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = new(Student)
		s.Name = "Jerry"
		s.Age = 18
		s.Class = "20005"
		s.Score = 100
	}
}

func DirectInvoke(s *Student) {
	s.Name = "Jerry"
	s.Age = 18
	s.Class = "20005"
	s.Score = 100
}

func InterfaceInvoke(i interface{}) {
	s := i.(*Student)
	s.Name = "Jerry"
	s.Age = 18
	s.Class = "20005"
	s.Score = 100
}

// BenchmarkDirectInvoke-12                1000000000               0.250 ns/op
func BenchmarkDirectInvoke(b *testing.B) {
	s := new(Student)
	for i := 0; i < b.N; i++ {
		DirectInvoke(s)
	}
	_ = s
}

// BenchmarkInterfaceInvoke-12             750168187                1.51 ns/op
func BenchmarkInterfaceInvoke(b *testing.B) {
	s := new(Student)
	for i := 0; i < b.N; i++ {
		InterfaceInvoke(s)
	}
	_ = s
}
