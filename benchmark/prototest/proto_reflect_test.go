package prototest

import (
	"testing"

	test "github.com/KippaZou/learn-reflection/benchmark/prototest/protoc-go"
	test2 "github.com/KippaZou/learn-reflection/benchmark/prototest/protoc-gogo-fast"
	proto2 "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

// BenchmarkProtoGoMarshal-12                       4792279               244 ns/op
func BenchmarkProtoGoMarshal(b *testing.B) {
	demo := &test.Demo{}
	_, err := proto2.Marshal(demo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(demo)
	}
}

// BenchmarkProtoGoUnMarshal-12                     6835482               173 ns/op
func BenchmarkProtoGoUnMarshal(b *testing.B) {
	demo := &test.Demo{}
	d, err := proto.Marshal(demo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(d, demo)
	}
}

// BenchmarkProtoGoGoFastMarshalV1-12              41996274                27.8 ns/op
func BenchmarkProtoGoGoFastMarshalV1(b *testing.B) {
	demo := &test2.Demo{}
	_, err := proto2.Marshal(demo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto2.Marshal(demo)
	}
}

// BenchmarkProtoGoGoFastUnMarshalV1-12            66546132                18.2 ns/op
func BenchmarkProtoGoGoFastUnMarshalV1(b *testing.B) {
	demo := &test2.Demo{}
	d, err := proto2.Marshal(demo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto2.Unmarshal(d, demo)
	}
}

// BenchmarkProtoGoGoFastMarshalV2-12              93389283                12.7 ns/op
func BenchmarkProtoGoGoFastMarshalV2(b *testing.B) {
	demo := &test2.Demo{}
	_, err := demo.Marshal()
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		demo.Marshal()
	}
}

// BenchmarkProtoGoGoFastUnMarshalV2-12            481613683                2.48 ns/op
func BenchmarkProtoGoGoFastUnMarshalV2(b *testing.B) {
	demo := &test2.Demo{}
	d, err := demo.Marshal()
	require.NoError(b, err)
	require.NoError(b, demo.Unmarshal(d))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		demo.Unmarshal(d)
	}
}
