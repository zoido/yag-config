package value_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zoido/yag-config/value"
)

func TestUint(t *testing.T) {
	// Given
	var num uint
	v := value.Uint(&num)

	// When
	err := v.Set("42")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(uint(42), num)
	r.Equal("42", v.String())
}

func TestUint_ParseError(t *testing.T) {
	// Given
	var num uint
	v := value.Uint(&num)

	// When
	err := v.Set("-42")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "-42")
}

func TestUint8(t *testing.T) {
	// Given
	var num uint8
	v := value.Uint8(&num)

	// When
	err := v.Set("42")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(uint8(42), num)
	r.Equal("42", v.String())
}

func TestUint8_ParseError(t *testing.T) {
	// Given
	var num uint8
	v := value.Uint8(&num)

	// When
	err := v.Set("-42")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "-42")
}

func TestUint8_ParseError_Overflow(t *testing.T) {
	// Given
	var num uint8
	v := value.Uint8(&num)

	// When
	err := v.Set("256")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}

func TestUint16(t *testing.T) {
	// Given
	var num uint16
	v := value.Uint16(&num)

	// When
	err := v.Set("60000")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(uint16(60000), num)
	r.Equal("60000", v.String())
}

func TestUint16_ParseError(t *testing.T) {
	// Given
	var num uint16
	v := value.Uint16(&num)

	// When
	err := v.Set("-42")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "-42")
}

func TestUint16_ParseError_Overflow(t *testing.T) {
	// Given
	var num uint16
	v := value.Uint16(&num)

	// When
	err := v.Set("128000")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}

func TestUint32(t *testing.T) {
	// Given
	var num uint32
	v := value.Uint32(&num)

	// When
	err := v.Set("4000000000")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(uint32(4_000_000_000), num)
	r.Equal("4000000000", v.String())
}

func TestUint32_ParseError(t *testing.T) {
	// Given
	var num uint32
	v := value.Uint32(&num)

	// When
	err := v.Set("-42")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "-42")
}

func TestUint32_ParseError_Overflow(t *testing.T) {
	// Given
	var num uint32
	v := value.Uint32(&num)

	// When
	err := v.Set("8000000000")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}

func TestUint64(t *testing.T) {
	// Given
	var num uint64
	v := value.Uint64(&num)

	// When
	err := v.Set("10000000000000000000")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(uint64(10_000_000_000_000_000_000), num)
	r.Equal("10000000000000000000", v.String())
}

func TestUint64_ParseError(t *testing.T) {
	// Given
	var num uint64
	v := value.Uint64(&num)

	// When
	err := v.Set("-42")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "-42")
}

func TestUint64_ParseError_Overflow(t *testing.T) {
	// Given
	var num uint64
	v := value.Uint64(&num)

	// When
	err := v.Set("20000000000000000000")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}
