package yag_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zoido/yag-config"
)

func TestArg_String(t *testing.T) {
	// Given
	var a, b string
	parser := yag.ArgParser{}
	parser.String(&a)
	parser.String(&b)

	// When
	err := parser.Parse([]string{"a", "b", "c", "d"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "a", a)
	require.Equal(t, "b", b)
}

func TestArg_Strings(t *testing.T) {
	// Given
	var (
		a string
		s []string
	)

	parser := yag.ArgParser{}
	parser.String(&a)
	parser.Strings(&s)

	// When
	err := parser.Parse([]string{"a", "b", "c", "d"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "a", a)
	require.Equal(t, []string{"b", "c", "d"}, s)
}

func TestArg_Int(t *testing.T) {
	// Given
	var a, b int
	parser := yag.ArgParser{}
	parser.Int(&a)
	parser.Int(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	require.NoError(t, err)
	require.Equal(t, 10, a)
	require.Equal(t, 20, b)
}

func TestArg_Ints(t *testing.T) {
	// Given
	var (
		s string
		i []int
	)

	parser := yag.ArgParser{}
	parser.String(&s)
	parser.Ints(&i)

	// When
	err := parser.Parse([]string{"a", "1", "2", "3"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "a", s)
	require.Equal(t, []int{1, 2, 3}, i)
}

func TestArg_Int8(t *testing.T) {
	// Given
	var a, b int8
	parser := yag.ArgParser{}
	parser.Int8(&a)
	parser.Int8(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	require.NoError(t, err)
	require.Equal(t, int8(10), a)
	require.Equal(t, int8(20), b)
}

func TestArg_Int8s(t *testing.T) {
	// Given
	var (
		s string
		i []int8
	)

	parser := yag.ArgParser{}
	parser.String(&s)
	parser.Int8s(&i)

	// When
	err := parser.Parse([]string{"a", "1", "2", "3"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "a", s)
	require.Equal(t, []int8{1, 2, 3}, i)
}

func TestArg_Int16(t *testing.T) {
	// Given
	var a, b int16
	parser := yag.ArgParser{}
	parser.Int16(&a)
	parser.Int16(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	require.NoError(t, err)
	require.Equal(t, int16(10), a)
	require.Equal(t, int16(20), b)
}

func TestArg_Int16s(t *testing.T) {
	// Given
	var (
		s string
		i []int16
	)

	parser := yag.ArgParser{}
	parser.String(&s)
	parser.Int16s(&i)

	// When
	err := parser.Parse([]string{"a", "1", "2", "3"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "a", s)
	require.Equal(t, []int16{1, 2, 3}, i)
}

func TestArg_Int32(t *testing.T) {
	// Given
	var a, b int32
	parser := yag.ArgParser{}
	parser.Int32(&a)
	parser.Int32(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	require.NoError(t, err)
	require.Equal(t, int32(10), a)
	require.Equal(t, int32(20), b)
}

func TestArg_Int32s(t *testing.T) {
	// Given
	var (
		s string
		i []int32
	)

	parser := yag.ArgParser{}
	parser.String(&s)
	parser.Int32s(&i)

	// When
	err := parser.Parse([]string{"a", "1", "2", "3"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "a", s)
	require.Equal(t, []int32{1, 2, 3}, i)
}

func TestArg_Int64(t *testing.T) {
	// Given
	var a, b int64
	parser := yag.ArgParser{}
	parser.Int64(&a)
	parser.Int64(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	require.NoError(t, err)
	require.Equal(t, int64(10), a)
	require.Equal(t, int64(20), b)
}

func TestArg_Int64s(t *testing.T) {
	// Given
	var (
		s string
		i []int64
	)

	parser := yag.ArgParser{}
	parser.String(&s)
	parser.Int64s(&i)

	// When
	err := parser.Parse([]string{"a", "1", "2", "3"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "a", s)
	require.Equal(t, []int64{1, 2, 3}, i)
}

func TestArg_Parse_Empty_Untouched(t *testing.T) {
	// Given
	var a string
	s := []string{"x", "y", "z"}

	parser := yag.ArgParser{}
	parser.String(&a)
	parser.Strings(&s)

	// When
	err := parser.Parse([]string{"a"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "a", a)
	require.Equal(t, []string{"x", "y", "z"}, s)
}

func TestArg_Parse_Error_SingleArgument(t *testing.T) {
	// Given
	var a, b int

	parser := yag.ArgParser{}
	parser.Int(&a)
	parser.Int(&b)

	// When
	err := parser.Parse([]string{"1", "x", "3", "4"})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "parsing argument on position 2")
}

func TestArg_Parse_Error_SingleArgument_WithName(t *testing.T) {
	// Given
	var a, b int

	parser := yag.ArgParser{}
	parser.Int(&a)
	parser.Int(&b, yag.WithName("b_name"))

	// When
	err := parser.Parse([]string{"1", "x", "3", "4"})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "parsing argument 'b_name' on position 2")
}

func TestArg_Parse_Error_MultipleArguments(t *testing.T) {
	// Given
	var (
		s string
		i []int
	)

	parser := yag.ArgParser{}
	parser.String(&s)
	parser.Ints(&i)

	// When
	err := parser.Parse([]string{"1", "2", "x", "4"})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "parsing int argument on position 3")
}

func TestArg_Parse_RequiredOption_FailsOnParse_WithPosition(t *testing.T) {
	// Given
	var a, b string

	parser := yag.ArgParser{}
	parser.String(&a)
	parser.String(&b, yag.Required())

	// When
	err := parser.Parse([]string{"a"})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "on position 2")
}

func TestArg_Parse_RequiredOption_FailsOnParse_WithName(t *testing.T) {
	// Given
	var a, b string

	parser := yag.ArgParser{}
	parser.String(&a)
	parser.String(&b, yag.Required(), yag.WithName("b_option"))

	// When
	err := parser.Parse([]string{"a"})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "on position 2")
	require.Contains(t, err.Error(), "'b_option' is required")
}
