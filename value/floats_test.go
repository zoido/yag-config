package value_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zoido/yag-config/value"
)

func TestFloat32(t *testing.T) {
	// Given
	var num float32
	v := value.Float32(&num)

	// When
	err := v.Set("3.14")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(float32(3.14), num)
	r.Equal("3.14", v.String())
}

func TestFloat32_ParseError(t *testing.T) {
	// Given
	var num float32
	v := value.Float32(&num)

	// When
	err := v.Set("3.14f")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "3.14")
}

func TestFloat32_ParseError_Overflow(t *testing.T) {
	// Given
	var num float32
	v := value.Float32(&num)

	// When
	err := v.Set("3e40")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}

func TestFloat64(t *testing.T) {
	// Given
	var num float64
	v := value.Float64(&num)

	// When
	err := v.Set("3.14")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(float64(3.14), num)
	r.Equal("3.14", v.String())
}

func TestFloat64_ParseError(t *testing.T) {
	// Given
	var num float64
	v := value.Float64(&num)

	// When
	err := v.Set("3.14f")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "3.14f")
}

func TestFloat64_ParseError_Overflow(t *testing.T) {
	// Given
	var num float64
	v := value.Float64(&num)

	// When
	err := v.Set("3e320")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "out of range")
}
