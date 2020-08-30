package value_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zoido/yag-config/value"
)

func TestInt(t *testing.T) {
	// Given
	var num int
	v := value.Int(&num)

	// When
	err := v.Set("-42")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(-42, num)
	r.Equal("-42", v.String())
}

func TestInt_ParseError(t *testing.T) {
	// Given
	var num int
	v := value.Int(&num)

	// When
	err := v.Set("3.14")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "3.14")
}

func TestInt8(t *testing.T) {
	// Given
	var num int8
	v := value.Int8(&num)

	// When
	err := v.Set("-42")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(int8(-42), num)
	r.Equal("-42", v.String())
}

func TestInt8_ParseError(t *testing.T) {
	// Given
	var num int8
	v := value.Int8(&num)

	// When
	err := v.Set("3.14")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "3.14")
}

func TestInt8_ParseError_Overflow(t *testing.T) {
	// Given
	var num int8
	v := value.Int8(&num)

	// When
	err := v.Set("-256")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}

func TestInt16(t *testing.T) {
	// Given
	var num int16
	v := value.Int16(&num)

	// When
	err := v.Set("-30000")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(int16(-30000), num)
	r.Equal("-30000", v.String())
}

func TestInt16_ParseError(t *testing.T) {
	// Given
	var num int16
	v := value.Int16(&num)

	// When
	err := v.Set("3.14")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "3.14")
}

func TestInt16_ParseError_Overflow(t *testing.T) {
	// Given
	var num int16
	v := value.Int16(&num)

	// When
	err := v.Set("64000")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}

func TestInt32(t *testing.T) {
	// Given
	var num int32
	v := value.Int32(&num)

	// When
	err := v.Set("-2000000000")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(int32(-2_000_000_000), num)
	r.Equal("-2000000000", v.String())
}

func TestInt32_ParseError(t *testing.T) {
	// Given
	var num int32
	v := value.Int32(&num)

	// When
	err := v.Set("3.14")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "3.14")
}

func TestInt32_ParseError_Overflow(t *testing.T) {
	// Given
	var num int32
	v := value.Int32(&num)

	// When
	err := v.Set("3000000000")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}

func TestInt64(t *testing.T) {
	// Given
	var num int64
	v := value.Int64(&num)

	// When
	err := v.Set("-9000000000000000000")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("-9000000000000000000", v.String())
}

func TestInt64_ParseError(t *testing.T) {
	// Given
	var num int64
	v := value.Int64(&num)

	// When
	err := v.Set("3.14")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "3.14")
}

func TestInt64_ParseError_Overflow(t *testing.T) {
	// Given
	var num int64
	v := value.Int64(&num)

	// When
	err := v.Set("10000000000000000000")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}
