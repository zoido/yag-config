package yag_test

import (
	"github.com/zoido/yag-config"
)

func (s *TypesTestSuite) TestUint() {
	// Given
	var num uint = 128

	y := yag.New()
	y.Uint(&num, "uint", "")

	// When
	err := y.Parse([]string{"-uint=42"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint(42), num)
}

func (s *TypesTestSuite) TestUint_DefaultValue() {
	// Given
	var num uint = 128

	y := yag.New()
	y.Uint(&num, "uint", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint(128), num)
}

func (s *TypesTestSuite) TestUint_ParseError() {
	// Given
	var num uint = 128

	y := yag.New()
	y.Uint(&num, "uint", "")

	// When
	err := y.Parse([]string{"-uint=3.14"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "3.14")
	s.Require().Contains(err.Error(), "-uint")
}

func (s *TypesTestSuite) TestUint8() {
	// Given
	var num uint8 = 42

	y := yag.New()
	y.Uint8(&num, "uint8", "")

	// When
	err := y.Parse([]string{"-uint8=42"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint8(42), num)
}

func (s *TypesTestSuite) TestUint8_DefaultValue() {
	// Given
	var num uint8 = 50

	y := yag.New()
	y.Uint8(&num, "uint", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint8(50), num)
}

func (s *TypesTestSuite) TestUint8_ParseError() {
	// Given
	var num uint8 = 42

	y := yag.New()
	y.Uint8(&num, "uint8", "")

	// When
	err := y.Parse([]string{"-uint8=-42"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "-42")
	s.Require().Contains(err.Error(), "-uint8")
}

func (s *TypesTestSuite) TestUint8_ParseError_Overflow() {
	// Given
	var num uint8 = 42

	y := yag.New()
	y.Uint8(&num, "uint8", "")

	// When
	err := y.Parse([]string{"-uint8=256"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-uint8")
}

func (s *TypesTestSuite) TestUint16() {
	// Given
	var num uint16 = 128

	y := yag.New()
	y.Uint16(&num, "uint16", "")

	// When
	err := y.Parse([]string{"-uint16=60000"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint16(60000), num)
}

func (s *TypesTestSuite) TestUint16_DefaultValue() {
	// Given
	var num uint16 = 128

	y := yag.New()
	y.Uint16(&num, "uint", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint16(128), num)
}

func (s *TypesTestSuite) TestUint16_ParseError() {
	// Given
	var num uint16 = 128

	y := yag.New()
	y.Uint16(&num, "uint16", "")

	// When
	err := y.Parse([]string{"-uint16=-128"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "-128")
	s.Require().Contains(err.Error(), "-uint16")
}

func (s *TypesTestSuite) TestUint16_ParseError_Overflow() {
	// Given
	var num uint16 = 128

	y := yag.New()
	y.Uint16(&num, "uint16", "")

	// When
	err := y.Parse([]string{"-uint16=128000"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-uint16")
}

func (s *TypesTestSuite) TestUint32() {
	// Given
	var num uint32 = 128

	y := yag.New()
	y.Uint32(&num, "uint32", "")

	// When
	err := y.Parse([]string{"-uint32=4000000000"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint32(4_000_000_000), num)
}

func (s *TypesTestSuite) TestUint32_DefaultValue() {
	// Given
	var num uint32 = 128

	y := yag.New()
	y.Uint32(&num, "uint", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint32(128), num)
}

func (s *TypesTestSuite) TestUint32_ParseError() {
	// Given
	var num uint32 = 128

	y := yag.New()
	y.Uint32(&num, "uint32", "")

	// When
	err := y.Parse([]string{"-uint32=-128"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "-128")
	s.Require().Contains(err.Error(), "-uint32")
}

func (s *TypesTestSuite) TestUint32_ParseError_Overflow() {
	// Given
	var num uint32 = 128

	y := yag.New()
	y.Uint32(&num, "uint32", "")

	// When
	err := y.Parse([]string{"-uint32=8000000000"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-uint32")
}

func (s *TypesTestSuite) TestUint64() {
	// Given
	var num uint64 = 128

	y := yag.New()
	y.Uint64(&num, "uint64", "")

	// When
	err := y.Parse([]string{"-uint64=10000000000000000000"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint64(10_000_000_000_000_000_000), num)
}

func (s *TypesTestSuite) TestUint64_DefaultValue() {
	// Given
	var num uint64 = 128

	y := yag.New()
	y.Uint64(&num, "uint", "")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal(uint64(128), num)
}

func (s *TypesTestSuite) TestUint64_ParseError() {
	// Given
	var num uint64 = 128

	y := yag.New()
	y.Uint64(&num, "uint64", "")

	// When
	err := y.Parse([]string{"-uint64=-128"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid value")
	s.Require().Contains(err.Error(), "-128")
	s.Require().Contains(err.Error(), "-uint64")
}

func (s *TypesTestSuite) TestUint64_ParseError_Overflow() {
	// Given
	var num uint64 = 128

	y := yag.New()
	y.Uint64(&num, "uint64", "")

	// When
	err := y.Parse([]string{"-uint64=20000000000000000000"})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "out of range")
	s.Require().Contains(err.Error(), "-uint64")
}
