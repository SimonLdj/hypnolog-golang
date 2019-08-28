package hypnolog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_General(t *testing.T) {
	res := false

	res = Create().Str("some string").Tag("test-example").Log()
	assert.True(t, res)
	res = Create().Set(struct{ValueA string; ValueB int}{"Hello", 77,}).Tag("test-example").Log()
	assert.True(t, res)
	res = Create().Set(map[string]string{"some-key":"some-value"}).Log()
	assert.True(t, res)
}

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
			LogString(tt.args.data)
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
		{name: "testWithString", args: args{data: "test string data"}},
		{name: "testWithCustomStruct", args: args{data: MyTestStruct{ParamA: "papapa", ParamB: 43}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogStruct(tt.args.data)
		})
	}
}

type MyTestStruct struct {
	ParamA string
	ParamB int
}

func TestSetHost(t *testing.T) {

	SetHost("not-existing-host")
	res := LogString("aaad")

	assert.Equal(t, false, res)

	SetHost("localhost")
	res = LogString("aaad")
	assert.Equal(t, true, res)
}

