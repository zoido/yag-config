package yag_test

import (
	"testing"
	"time"

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
	y.String(&str, "string", "")

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
	y.String(&str, "string", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("default value", str)
}

func (s *TypesTestSuite) TestBool_BoolFlag() {
	// Given
	var b bool

	y := yag.New()
	y.Bool(&b, "bool", "")

	// When
	err := y.Parse([]string{"-bool"})

	// Then
	s.Require().NoError(err)
	s.Require().True(b)
}

func (s *TypesTestSuite) TestBool_FlagWithValue() {
	// Given
	var b bool

	y := yag.New()
	y.Bool(&b, "bool", "")

	// When
	err := y.Parse([]string{"-bool=true"})

	// Then
	s.Require().NoError(err)
	s.Require().True(b)
}

func (s *TypesTestSuite) TestBool_DefaultValue() {
	// Given
	b := true

	y := yag.New()
	y.Bool(&b, "bool", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().True(b)
}

func (s *TypesTestSuite) TestBool_ParseError() {
	// Given
	var b bool

	y := yag.New()
	y.Bool(&b, "bool", "")

	// When
	err := y.Parse([]string{"-bool=yes"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid boolean value")
	s.Require().Contains(err.Error(), "yes")
	s.Require().Contains(err.Error(), "-bool")
}

func (s *TypesTestSuite) TestDuration() {
	// Given
	var dur time.Duration

	y := yag.New()
	y.Duration(&dur, "dur", "")

	// When
	err := y.Parse([]string{"-dur=10s"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(float64(10), dur.Seconds())
}

func (s *TypesTestSuite) TestDuration_DefaultValue() {
	// Given
	dur := time.Minute

	y := yag.New()
	y.Duration(&dur, "int", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(float64(60), dur.Seconds())
}

func (s *TypesTestSuite) TestDuration_ParseError() {
	// Given
	var dur time.Duration

	y := yag.New()
	y.Duration(&dur, "dur", "")

	// When
	err := y.Parse([]string{"-dur=10x"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "10x")
	s.Require().Contains(err.Error(), "-dur")
}

type testFlagValue struct{}

func (tfv *testFlagValue) Set(val string) error {
	return nil
}

func (tfv *testFlagValue) String() string {
	return "test_flag_value"
}

func (s *TypesTestSuite) TestValue() {
	// Given
	val := &testFlagValue{}

	y := yag.New()
	y.Value(val, "val", "")

	// When
	err := y.Parse([]string{"-val=test_string"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("test_flag_value", val.String())
}
