package yag_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/zoido/yag-config"
)

type ArgParserTestSuite struct {
	suite.Suite
}

func TestArgParser(t *testing.T) {
	suite.Run(t, new(ArgParserTestSuite))
}

func (ts *ArgParserTestSuite) TestString() {
	// Given
	var a, b string
	parser := yag.ArgParser{}
	parser.String(&a)
	parser.String(&b)

	// When
	err := parser.Parse([]string{"a", "b", "c", "d"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("a", a)
	ts.Require().Equal("b", b)
}

func (ts *ArgParserTestSuite) TestStrings() {
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
	ts.Require().NoError(err)
	ts.Require().Equal("a", a)
	ts.Require().Equal([]string{"b", "c", "d"}, s)
}

func (ts *ArgParserTestSuite) TestInt() {
	// Given
	var a, b int
	parser := yag.ArgParser{}
	parser.Int(&a)
	parser.Int(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(10, a)
	ts.Require().Equal(20, b)
}

func (ts *ArgParserTestSuite) TestInts() {
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
	ts.Require().NoError(err)
	ts.Require().Equal("a", s)
	ts.Require().Equal([]int{1, 2, 3}, i)
}

func (ts *ArgParserTestSuite) TestInt8() {
	// Given
	var a, b int8
	parser := yag.ArgParser{}
	parser.Int8(&a)
	parser.Int8(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(int8(10), a)
	ts.Require().Equal(int8(20), b)
}

func (ts *ArgParserTestSuite) TestInt8s() {
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
	ts.Require().NoError(err)
	ts.Require().Equal("a", s)
	ts.Require().Equal([]int8{1, 2, 3}, i)
}

func (ts *ArgParserTestSuite) TestInt16() {
	// Given
	var a, b int16
	parser := yag.ArgParser{}
	parser.Int16(&a)
	parser.Int16(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(int16(10), a)
	ts.Require().Equal(int16(20), b)
}

func (ts *ArgParserTestSuite) TestInt16s() {
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
	ts.Require().NoError(err)
	ts.Require().Equal("a", s)
	ts.Require().Equal([]int16{1, 2, 3}, i)
}

func (ts *ArgParserTestSuite) TestInt32() {
	// Given
	var a, b int32
	parser := yag.ArgParser{}
	parser.Int32(&a)
	parser.Int32(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(int32(10), a)
	ts.Require().Equal(int32(20), b)
}

func (ts *ArgParserTestSuite) TestInt32s() {
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
	ts.Require().NoError(err)
	ts.Require().Equal("a", s)
	ts.Require().Equal([]int32{1, 2, 3}, i)
}

func (ts *ArgParserTestSuite) TestInt64() {
	// Given
	var a, b int64
	parser := yag.ArgParser{}
	parser.Int64(&a)
	parser.Int64(&b)

	// When
	err := parser.Parse([]string{"10", "20"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(int64(10), a)
	ts.Require().Equal(int64(20), b)
}

func (ts *ArgParserTestSuite) TestInt64s() {
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
	ts.Require().NoError(err)
	ts.Require().Equal("a", s)
	ts.Require().Equal([]int64{1, 2, 3}, i)
}

func (ts *ArgParserTestSuite) TestParse_Empty_Untouched() {
	// Given
	var a string
	s := []string{"x", "y", "z"}

	parser := yag.ArgParser{}
	parser.String(&a)
	parser.Strings(&s)

	// When
	err := parser.Parse([]string{"a"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("a", a)
	ts.Require().Equal([]string{"x", "y", "z"}, s)
}

func (ts *ArgParserTestSuite) TestParse_RequiredOption_FailsOnParse_WithPosition() {
	// Given
	var a, b string

	parser := yag.ArgParser{}
	parser.String(&a)
	parser.String(&b, yag.Required())

	// When
	err := parser.Parse([]string{"a"})

	// Then
	ts.Require().Error(err)
	ts.Require().Contains(err.Error(), "on position 2")
}

func (ts *ArgParserTestSuite) TestParse_RequiredOption_FailsOnParse_WithName() {
	// Given
	var a, b string

	parser := yag.ArgParser{}
	parser.String(&a)
	parser.String(&b, yag.Required(), yag.WithName("b_option"))

	// When
	err := parser.Parse([]string{"a"})

	// Then
	ts.Require().Error(err)
	ts.Require().Contains(err.Error(), "on position 2")
	ts.Require().Contains(err.Error(), "'b_option' is required")
}
