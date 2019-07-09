package hypnolog

import (
	"testing"
)

func TestHypnoLogString(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testA", args: args{data: "test data"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HypnoLogString(tt.args.data)
		})
	}
}

func TestHypnoLogStruct(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testWithString", args: args{data: "test data"}},
		{name: "testWithCustomStruct", args: args{data: MyTestStruct{ParamA:"papapa", ParamB:43}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HypnoLogStruct(tt.args.data)
		})
	}
}

type MyTestStruct struct {
	ParamA string
	ParamB int
}
