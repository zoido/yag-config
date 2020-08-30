package yag_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zoido/yag-config"
)

type testFlagValue struct{}

func (tfv *testFlagValue) Set(val string) error {
	return nil
}

func (tfv *testFlagValue) String() string {
	return "test_flag_value"
}

func TestValue(t *testing.T) {
	// Given
	val := &testFlagValue{}

	y := yag.New()
	y.Value(val, "val", "")

	// When
	err := y.Parse([]string{"-val=test_string"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("test_flag_value", val.String())
}
