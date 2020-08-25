package yag_test

import (
	"github.com/zoido/yag-config"
)

func (s *TypesTestSuite) TestFloat32() {
	// Given
	var num float32 = 6.626e-34

	y := yag.New()
	y.Float32(&num, "float32", "")

	// When
	err := y.Parse([]string{"-float32=3.14"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(float32(3.14), num)
}

func (s *TypesTestSuite) TestFloat32_DefaultValue() {
	// Given
	var num float32 = 2.72

	y := yag.New()
	y.Float32(&num, "int", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(float32(2.72), num)
}

func (s *TypesTestSuite) TestFloat32_ParseError() {
	// Given
	var num float32 = 6.626e-34

	y := yag.New()
	y.Float32(&num, "float32", "")

	// When
	err := y.Parse([]string{"-float32=3.14f"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "3.14")
	s.Require().Contains(err.Error(), "-float32")
}

func (s *TypesTestSuite) TestFloat32_ParseError_Overflow() {
	// Given
	var num float32 = 6.626e-34

	y := yag.New()
	y.Float32(&num, "float32", "")

	// When
	err := y.Parse([]string{"-float32=3e40"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-float32")
}
