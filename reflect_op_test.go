// referred from https://colobu.com/2019/01/29/go-reflect-performance/
package test

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

func BenchmarkDirect_New(b *testing.B) {
	var s *Student
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = new(Student)
	}
	_ = s
}

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

func BenchmarkDirectInvoke(b *testing.B) {
	s := new(Student)
	for i := 0; i < b.N; i++ {
		DirectInvoke(s)
	}
	_ = s
}

func BenchmarkInterfaceInvoke(b *testing.B) {
	s := new(Student)
	for i := 0; i < b.N; i++ {
		InterfaceInvoke(s)
	}
	_ = s
}
