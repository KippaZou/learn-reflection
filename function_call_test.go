package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func test(_ string, _ []int) {}

func TestBatchCall(t *testing.T) {
	err := BatchCall(2, test, "", []int{1, 2})
	require.NoError(t, err)
	err = BatchCall(2, test, []int{1, 2, 3})
	require.Error(t, err)
}

// BenchmarkBatch-12    	 2363384	       464 ns/op
func BenchmarkBatchCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BatchCall(2, test, "", []int{1, 2, 3, 4})
	}
}

// BenchmarkCall-12    	1000000000	         0.518 ns/op
func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test("", []int{1, 2})
		test("", []int{3, 4})
	}
}
