package yag_test

import (
	"github.com/zoido/yag-config"
)

func (s *TypesTestSuite) TestInt() {
	// Given
	num := 128

	y := yag.New()
	y.Int(&num, "int", "")

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
	y.Int(&num, "int", "")

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
	y.Int(&num, "int", "")

	// When
	err := y.Parse([]string{"-int=3.14"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "3.14")
	s.Require().Contains(err.Error(), "-int")
}

func (s *TypesTestSuite) TestInt8() {
	// Given
	var num int8 = 4

	y := yag.New()
	y.Int8(&num, "int8", "")

	// When
	err := y.Parse([]string{"-int8=-42"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(int8(-42), num)
}

func (s *TypesTestSuite) TestInt8_DefaultValue() {
	// Given
	var num int8 = 42

	y := yag.New()
	y.Int8(&num, "int", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(int8(42), num)
}

func (s *TypesTestSuite) TestInt8_ParseError() {
	// Given
	var num int8 = 42

	y := yag.New()
	y.Int8(&num, "int8", "")

	// When
	err := y.Parse([]string{"-int8=3.14"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "3.14")
	s.Require().Contains(err.Error(), "-int8")
}

func (s *TypesTestSuite) TestInt8_ParseError_Overflow() {
	// Given
	var num int8 = 42

	y := yag.New()
	y.Int8(&num, "int8", "")

	// When
	err := y.Parse([]string{"-int8=-256"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-int8")
}

func (s *TypesTestSuite) TestInt16() {
	// Given
	var num int16 = 128

	y := yag.New()
	y.Int16(&num, "int16", "")

	// When
	err := y.Parse([]string{"-int16=-30000"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(int16(-30000), num)
}

func (s *TypesTestSuite) TestInt16_DefaultValue() {
	// Given
	var num int16 = 128

	y := yag.New()
	y.Int16(&num, "int", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(int16(128), num)
}

func (s *TypesTestSuite) TestInt16_ParseError() {
	// Given
	var num int16 = 128

	y := yag.New()
	y.Int16(&num, "int16", "")

	// When
	err := y.Parse([]string{"-int16=3.14"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "3.14")
	s.Require().Contains(err.Error(), "-int16")
}

func (s *TypesTestSuite) TestInt16_ParseError_Overflow() {
	// Given
	var num int16 = 128

	y := yag.New()
	y.Int16(&num, "int16", "")

	// When
	err := y.Parse([]string{"-int16=-64000"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-int16")
}

func (s *TypesTestSuite) TestInt32() {
	// Given
	var num int32 = 128

	y := yag.New()
	y.Int32(&num, "int32", "")

	// When
	err := y.Parse([]string{"-int32=2000000000"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(int32(2_000_000_000), num)
}

func (s *TypesTestSuite) TestInt32_DefaultValue() {
	// Given
	var num int32 = 128

	y := yag.New()
	y.Int32(&num, "int", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(int32(128), num)
}

func (s *TypesTestSuite) TestInt32_ParseError() {
	// Given
	var num int32 = 128

	y := yag.New()
	y.Int32(&num, "int32", "")

	// When
	err := y.Parse([]string{"-int32=3.14"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "3.14")
	s.Require().Contains(err.Error(), "-int32")
}

func (s *TypesTestSuite) TestInt32_ParseError_Overflow() {
	// Given
	var num int32 = 128

	y := yag.New()
	y.Int32(&num, "int32", "")

	// When
	err := y.Parse([]string{"-int32=3000000000"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-int32")
}
func (s *TypesTestSuite) TestInt64() {
	// Given
	var num int64 = 128

	y := yag.New()
	y.Int64(&num, "int64", "")

	// When
	err := y.Parse([]string{"-int64=9000000000000000000"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(int64(9_000_000_000_000_000_000), num)
}

func (s *TypesTestSuite) TestInt64_DefaultValue() {
	// Given
	var num int64 = 128

	y := yag.New()
	y.Int64(&num, "int", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(int64(128), num)
}

func (s *TypesTestSuite) TestInt64_ParseError() {
	// Given
	var num int64 = 128

	y := yag.New()
	y.Int64(&num, "int64", "")

	// When
	err := y.Parse([]string{"-int64=3.14"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "3.14")
	s.Require().Contains(err.Error(), "-int64")
}

func (s *TypesTestSuite) TestInt64_ParseError_Overflow() {
	// Given
	var num int64 = 128

	y := yag.New()
	y.Int64(&num, "int64", "")

	// When
	err := y.Parse([]string{"-int64=10000000000000000000"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-int64")
}
