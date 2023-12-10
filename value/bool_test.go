package value_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zoido/yag-config/value"
)

func TestBool(t *testing.T) {
	// Given
	var b bool
	v := value.Bool(&b)

	// When
	err := v.Set("true")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.True(b)
	r.Equal("true", v.String())
}

func TestBool_ParseError(t *testing.T) {
	// Given
	var b bool
	v := value.Bool(&b)

	// When
	err := v.Set("yes")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "yes")
}

func TestBool_IsBooleanFlag(t *testing.T) {
	// Given
	var b bool

	// When
	v := value.Bool(&b)

	// Then
	require.True(t, value.IsBoolFlag(v))
}

func TestIsBooleanFlag_NotBoolFlag(t *testing.T) {
	// When
	result := value.IsBoolFlag(&normalFlagValue{})

	// Then
	require.False(t, result)
}

func TestIsBooleanFlag_BoolFlagFalse(t *testing.T) {
	// When
	result := value.IsBoolFlag(&booleanFlagReturningFalse{})

	// Then
	require.False(t, result)
}

func TestIsBooleanFlag_BoolFlagTrue(t *testing.T) {
	// When
	result := value.IsBoolFlag(&booleanFlagReturningTrue{})

	// Then
	require.True(t, result)
}

type normalFlagValue struct{}

func (*normalFlagValue) Set(_ string) error {
	return nil
}

func (*normalFlagValue) String() string {
	return ""
}

type booleanFlagReturningFalse struct {
	normalFlagValue
}

func (*booleanFlagReturningFalse) IsBoolFlag() bool {
	return false
}

type booleanFlagReturningTrue struct {
	normalFlagValue
}

func (*booleanFlagReturningTrue) IsBoolFlag() bool {
	return true
}
