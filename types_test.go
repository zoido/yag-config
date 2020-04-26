package yag_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/zoido/yag-config"
)

type TypesTestSuite struct {
	suite.Suite
}

func TestTypesTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}

func (s *TypesTestSuite) TestString() {
	// Given
	str := "default value"

	y := yag.New()
	y.Register(&str, "string", "")

	// When
	err := y.Parse([]string{"-string=value"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("value", str)
}

func (s *TypesTestSuite) TestString_DefaultValue() {
	// Given
	str := "default value"

	y := yag.New()
	y.Register(&str, "string", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("default value", str)
}

func (s *TypesTestSuite) TestInt() {
	// Given
	num := 128

	y := yag.New()
	y.Register(&num, "int", "")

	// When
	err := y.Parse([]string{"-int=42"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(42, num)
}

func (s *TypesTestSuite) TestInt_DefaultValue() {
	// Given
	num := 128

	y := yag.New()
	y.Register(&num, "int", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(128, num)
}

func (s *TypesTestSuite) TestInt_ParseError() {
	// Given
	num := 128

	y := yag.New()
	y.Register(&num, "int", "")

	// When
	err := y.Parse([]string{"-int=3.14"})

	// Then
	s.Require().Error(err)
	// s.Require().Contains(err.Error(), "invalid value")
	// s.Require().Contains(err.Error(), "3.14")
	// s.Require().Contains(err.Error(), "-int")
}
