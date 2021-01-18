package test

import (
	"encoding/json"
	"sort"
	"testing"
)

func BenchmarkJSONReflectionMarshal(b *testing.B) {
	s := &Struct{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		json.Marshal(s)
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	s := &Struct{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.MarshalJSON()
	}
}

func BenchmarkJSONReflectionUnmarshal(b *testing.B) {
	s := &Struct{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
	}
	data, _ := json.Marshal(s)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &s)
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	s := &Struct{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
	}
	data, _ := json.Marshal(s)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.UnmarshalJSON(data)
	}
}

func BenchmarkJSONReflectionMarshalIface(b *testing.B) {
	var s Iface = &Struct{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		json.Marshal(s)
	}
}

func BenchmarkJSONReflectionUnmarshalIface(b *testing.B) {
	var s Iface = &Struct{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
	}
	data, _ := json.Marshal(s)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &s)
	}
}

// BenchmarkStructMethodCall-12  1000000000  0.252 ns/op
func BenchmarkStructMethodCall(b *testing.B) {
	s := &Struct{}
	for i := 0; i < b.N; i++ {
		s.Foo()
	}
}

// BenchmarkIfaceMethodCall-12  952953033  1.26 ns/op
func BenchmarkIfaceMethodCall(b *testing.B) {
	var s Iface = &Struct{}
	for i := 0; i < b.N; i++ {
		s.Foo()
	}
}

type SortableIface interface {
	Number() int
}

type Sortable struct {
	number int
}

func (s Sortable) Number() int {
	return s.number
}

type SortableIfaceByNumber []SortableIface

func (a SortableIfaceByNumber) Len() int           { return len(a) }
func (a SortableIfaceByNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableIfaceByNumber) Less(i, j int) bool { return a[i].Number() < a[j].Number() }

type SortableByNumber []Sortable

func (a SortableByNumber) Len() int           { return len(a) }
func (a SortableByNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableByNumber) Less(i, j int) bool { return a[i].Number() < a[j].Number() }

// BenchmarkSortStruct-12                18          63566411 ns/op
func BenchmarkSortStruct(b *testing.B) {
	s := make(SortableByNumber, 1000000)
	for i := 0; i < 1000000; i++ {
		s[i] = Sortable{i}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Sort(s)
	}
}

// BenchmarkSortIface-12                  6         175996484 ns/op
func BenchmarkSortIface(b *testing.B) {
	s := make(SortableIfaceByNumber, 1000000)
	for i := 0; i < 1000000; i++ {
		s[i] = Sortable{i}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Sort(s)
	}
}
