package prototest

import (
	"testing"

	test "github.com/KippaZou/learn-reflection/prototest/protoc-go"
	test2 "github.com/KippaZou/learn-reflection/prototest/protoc-gogo-fast"
	proto2 "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func BenchmarkProtoGoMarshal(b *testing.B) {
	demo := &test.Demo{}
	_, err := proto2.Marshal(demo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(demo)
	}
}

func BenchmarkProtoGoUnMarshal(b *testing.B) {
	demo := &test.Demo{}
	d, err := proto.Marshal(demo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(d, demo)
	}
}

func BenchmarkProtoGoGoFastMarshalV1(b *testing.B) {
	demo := &test2.Demo{}
	_, err := proto2.Marshal(demo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto2.Marshal(demo)
	}
}

func BenchmarkProtoGoGoFastUnMarshalV1(b *testing.B) {
	demo := &test2.Demo{}
	d, err := proto2.Marshal(demo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto2.Unmarshal(d, demo)
	}
}

func BenchmarkProtoGoGoFastMarshalV2(b *testing.B) {
	demo := &test2.Demo{}
	_, err := demo.Marshal()
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		demo.Marshal()
	}
}

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
