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
	var str string

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
