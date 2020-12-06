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
