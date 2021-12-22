package args_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zoido/yag-config/args"
)

func TestInt(t *testing.T) {
	// Given
	var i []int
	parser := args.Ints(&i)

	// When
	n, err := parser.Parse([]string{"1", "2", "3", "4"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(4, n)
	r.Equal([]int{1, 2, 3, 4}, i)
}

func TestInt_ParseError(t *testing.T) {
	// Given
	var i []int
	parser := args.Ints(&i)

	// When
	n, err := parser.Parse([]string{"1", "3.14"})

	// Then
	r := require.New(t)
	r.Error(err)
	r.Equal(1, n)
	r.Contains(err.Error(), "3.14")
}

func TestInts_Empty(t *testing.T) {
	// Given
	i := []int{10, 11, 12}
	parser := args.Ints(&i)

	// When
	n, err := parser.Parse([]string{})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(0, n)
	r.Equal([]int{}, i)
}

func TestInt8(t *testing.T) {
	// Given
	var i []int8
	parser := args.Int8s(&i)

	// When
	n, err := parser.Parse([]string{"1", "2", "3", "4"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(4, n)
	r.Equal([]int8{1, 2, 3, 4}, i)
}

func TestInt8_ParseError(t *testing.T) {
	// Given
	var i []int8
	parser := args.Int8s(&i)

	// When
	n, err := parser.Parse([]string{"1", "3.14"})

	// Then
	r := require.New(t)
	r.Error(err)
	r.Equal(1, n)
	r.Contains(err.Error(), "3.14")
}

func TestInt8s_Empty(t *testing.T) {
	// Given
	i := []int8{10, 11, 12}
	parser := args.Int8s(&i)

	// When
	n, err := parser.Parse([]string{})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(0, n)
	r.Equal([]int8{}, i)
}

func TestInt16(t *testing.T) {
	// Given
	var i []int16
	parser := args.Int16s(&i)

	// When
	n, err := parser.Parse([]string{"1", "2", "3", "4"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(4, n)
	r.Equal([]int16{1, 2, 3, 4}, i)
}

func TestInt16_ParseError(t *testing.T) {
	// Given
	var i []int16
	parser := args.Int16s(&i)

	// When
	n, err := parser.Parse([]string{"1", "3.14"})

	// Then
	r := require.New(t)
	r.Error(err)
	r.Equal(1, n)
	r.Contains(err.Error(), "3.14")
}

func TestInt16s_Empty(t *testing.T) {
	// Given
	i := []int16{10, 11, 12}
	parser := args.Int16s(&i)

	// When
	n, err := parser.Parse([]string{})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(0, n)
	r.Equal([]int16{}, i)
}

func TestInt32(t *testing.T) {
	// Given
	var i []int32
	parser := args.Int32s(&i)

	// When
	n, err := parser.Parse([]string{"1", "2", "3", "4"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(4, n)
	r.Equal([]int32{1, 2, 3, 4}, i)
}

func TestInt32_ParseError(t *testing.T) {
	// Given
	var i []int32
	parser := args.Int32s(&i)

	// When
	n, err := parser.Parse([]string{"1", "3.14"})

	// Then
	r := require.New(t)
	r.Error(err)
	r.Equal(1, n)
	r.Contains(err.Error(), "3.14")
}

func TestInt32s_Empty(t *testing.T) {
	// Given
	i := []int32{10, 11, 12}
	parser := args.Int32s(&i)

	// When
	n, err := parser.Parse([]string{})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(0, n)
	r.Equal([]int32{}, i)
}

func TestInt64(t *testing.T) {
	// Given
	var i []int64
	parser := args.Int64s(&i)

	// When
	n, err := parser.Parse([]string{"1", "2", "3", "4"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(4, n)
	r.Equal([]int64{1, 2, 3, 4}, i)
}

func TestInt64_ParseError(t *testing.T) {
	// Given
	var i []int64
	parser := args.Int64s(&i)

	// When
	n, err := parser.Parse([]string{"1", "3.14"})

	// Then
	r := require.New(t)
	r.Error(err)
	r.Equal(1, n)
	r.Contains(err.Error(), "3.14")
}

func TestInt64s_Empty(t *testing.T) {
	// Given
	i := []int64{10, 11, 12}
	parser := args.Int64s(&i)

	// When
	n, err := parser.Parse([]string{})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(0, n)
	r.Equal([]int64{}, i)
}
