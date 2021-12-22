package yag_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zoido/yag-config"
)

func TestArgParser_String(t *testing.T) {
	// Given
	var a, b string
	y := yag.New()
	y.Args().String(&a)
	y.Args().String(&b)

	// When
	err := y.Parse([]string{"a", "b", "c", "d"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("a", a)
	r.Equal("b", b)
}

func TestArgParser_Strings(t *testing.T) {
	// Given
	var (
		a string
		s []string
	)

	y := yag.New()
	y.Args().String(&a)
	y.Args().Strings(&s)

	// When
	err := y.Parse([]string{"a", "b", "c", "d"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("a", a)
	r.Equal([]string{"b", "c", "d"}, s)
}

func TestArgParser_Int(t *testing.T) {
	// Given
	var a, b int
	y := yag.New()
	y.Args().Int(&a)
	y.Args().Int(&b)

	// When
	err := y.Parse([]string{"10", "20"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(10, a)
	r.Equal(20, b)
}

func TestArgParser_Ints(t *testing.T) {
	// Given
	var (
		s string
		i []int
	)

	y := yag.New()
	y.Args().String(&s)
	y.Args().Ints(&i)

	// When
	err := y.Parse([]string{"a", "1", "2", "3"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("a", s)
	r.Equal([]int{1, 2, 3}, i)
}

func TestArgParser_Int8(t *testing.T) {
	// Given
	var a, b int8
	y := yag.New()
	y.Args().Int8(&a)
	y.Args().Int8(&b)

	// When
	err := y.Parse([]string{"10", "20"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(int8(10), a)
	r.Equal(int8(20), b)
}

func TestArgParser_Int8s(t *testing.T) {
	// Given
	var (
		s string
		i []int8
	)

	y := yag.New()
	y.Args().String(&s)
	y.Args().Int8s(&i)

	// When
	err := y.Parse([]string{"a", "1", "2", "3"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("a", s)
	r.Equal([]int8{1, 2, 3}, i)
}

func TestArgParser_Int16(t *testing.T) {
	// Given
	var a, b int16
	y := yag.New()
	y.Args().Int16(&a)
	y.Args().Int16(&b)

	// When
	err := y.Parse([]string{"10", "20"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(int16(10), a)
	r.Equal(int16(20), b)
}

func TestArgParser_Int16s(t *testing.T) {
	// Given
	var (
		s string
		i []int16
	)

	y := yag.New()
	y.Args().String(&s)
	y.Args().Int16s(&i)

	// When
	err := y.Parse([]string{"a", "1", "2", "3"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("a", s)
	r.Equal([]int16{1, 2, 3}, i)
}

func TestArgParser_Int32(t *testing.T) {
	// Given
	var a, b int32
	y := yag.New()
	y.Args().Int32(&a)
	y.Args().Int32(&b)

	// When
	err := y.Parse([]string{"10", "20"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(int32(10), a)
	r.Equal(int32(20), b)
}

func TestArgParser_Int32s(t *testing.T) {
	// Given
	var (
		s string
		i []int32
	)

	y := yag.New()
	y.Args().String(&s)
	y.Args().Int32s(&i)

	// When
	err := y.Parse([]string{"a", "1", "2", "3"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("a", s)
	r.Equal([]int32{1, 2, 3}, i)
}

func TestArgParser_Int64(t *testing.T) {
	// Given
	var a, b int64
	y := yag.New()
	y.Args().Int64(&a)
	y.Args().Int64(&b)

	// When
	err := y.Parse([]string{"10", "20"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(int64(10), a)
	r.Equal(int64(20), b)
}

func TestArgParser_Int64s(t *testing.T) {
	// Given
	var (
		s string
		i []int64
	)

	y := yag.New()
	y.Args().String(&s)
	y.Args().Int64s(&i)

	// When
	err := y.Parse([]string{"a", "1", "2", "3"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("a", s)
	r.Equal([]int64{1, 2, 3}, i)
}

func TestArgParser_Parse_Empty_Untouched(t *testing.T) {
	// Given
	var a string
	s := []string{"x", "y", "z"}

	y := yag.New()
	y.Args().String(&a)
	y.Args().Strings(&s)

	// When
	err := y.Parse([]string{"a"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("a", a)
	r.Equal([]string{"x", "y", "z"}, s)
}

func TestArgParser_Parse_Error_SingleArgument(t *testing.T) {
	// Given
	var a, b int

	y := yag.New()
	y.Args().Int(&a)
	y.Args().Int(&b)

	// When
	err := y.Parse([]string{"1", "x", "3", "4"})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "parsing argument on position 2")
}

func TestArgParser_Parse_Error_SingleArgument_WithName(t *testing.T) {
	// Given
	var a, b int

	y := yag.New()
	y.Args().Int(&a)
	y.Args().Int(&b, yag.WithName("b_name"))

	// When
	err := y.Parse([]string{"1", "x", "3", "4"})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "parsing argument 'b_name' on position 2")
}

func TestArgParser_Parse_Error_MultipleArguments(t *testing.T) {
	// Given
	var (
		s string
		i []int
	)

	y := yag.New()
	y.Args().String(&s)
	y.Args().Ints(&i)

	// When
	err := y.Parse([]string{"1", "2", "3.14", "4"})

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "on position 3")
	r.Contains(err.Error(), "3.14")
}

func TestArgParser_Parse_RequiredOption_FailsOnParse_WithPosition(t *testing.T) {
	// Given
	var a, b string

	y := yag.New()
	y.Args().String(&a)
	y.Args().String(&b, yag.Required())

	// When
	err := y.Parse([]string{"a"})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "on position 2")
}

func TestArgParser_Parse_RequiredOption_FailsOnParse_WithName(t *testing.T) {
	// Given
	var a, b string

	y := yag.New()
	y.Args().String(&a)
	y.Args().String(&b, yag.Required(), yag.WithName("b_option"))

	// When
	err := y.Parse([]string{"a"})

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "on position 2")
	r.Contains(err.Error(), "\"b_option\" is required")
}
