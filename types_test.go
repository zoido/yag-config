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

func TestTypes(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}

func (ts *TypesTestSuite) TestString() {
	// Given
	var s string

	y := yag.New()
	y.String(&s, "val", "")

	// When
	err := y.Parse([]string{"-val=test_string"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("test_string", s)
}

func (ts *TypesTestSuite) TestInt() {
	// Given
	var n int

	y := yag.New()
	y.Int(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(8, n)
}

func (ts *TypesTestSuite) TestInt8() {
	// Given
	var n int8

	y := yag.New()
	y.Int8(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(int8(8), n)
}

func (ts *TypesTestSuite) TestInt16() {
	// Given
	var n int16

	y := yag.New()
	y.Int16(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(int16(8), n)
}

func (ts *TypesTestSuite) TestInt32() {
	// Given
	var n int32

	y := yag.New()
	y.Int32(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(int32(8), n)
}

func (ts *TypesTestSuite) TestInt64() {
	// Given
	var n int64

	y := yag.New()
	y.Int64(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(int64(8), n)
}

func (ts *TypesTestSuite) TestUint() {
	// Given
	var n uint

	y := yag.New()
	y.Uint(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(uint(8), n)
}

func (ts *TypesTestSuite) TestUint8() {
	// Given
	var n uint8

	y := yag.New()
	y.Uint8(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(uint8(8), n)
}

func (ts *TypesTestSuite) TestUint16() {
	// Given
	var n uint16

	y := yag.New()
	y.Uint16(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(uint16(8), n)
}

func (ts *TypesTestSuite) TestUint32() {
	// Given
	var n uint32

	y := yag.New()
	y.Uint32(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(uint32(8), n)
}

func (ts *TypesTestSuite) TestUint64() {
	// Given
	var n uint64

	y := yag.New()
	y.Uint64(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(uint64(8), n)
}

func (ts *TypesTestSuite) TestFloat32() {
	// Given
	var n float32

	y := yag.New()
	y.Float32(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=6.626E-34"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(float32(6.626e-34), n)
}

func (ts *TypesTestSuite) TestFloat64() {
	// Given
	var n float64

	y := yag.New()
	y.Float64(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=6.626E-34"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(float64(6.626e-34), n)
}

func (ts *TypesTestSuite) TestDuration() {
	// Given
	var d time.Duration

	y := yag.New()
	y.Duration(&d, "val", "")

	// When
	err := y.Parse([]string{"-val=10h30m15s"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal(float64(10*60*60+30*60+15), d.Seconds())
}

func (ts *TypesTestSuite) TestBool() {
	// Given
	var b bool

	y := yag.New()
	y.Bool(&b, "val", "")

	// When
	err := y.Parse([]string{"-val=true"})

	// Then
	ts.Require().NoError(err)
	ts.Require().True(b)
}

func (ts *TypesTestSuite) TestBool_BoolFlag() {
	// Given
	var b bool

	y := yag.New()
	y.Bool(&b, "val", "")

	// When
	err := y.Parse([]string{"-val"})

	// Then
	ts.Require().NoError(err)
	ts.Require().True(b)
}

func (ts *TypesTestSuite) TestBool_False() {
	// Given
	b := false

	y := yag.New()
	y.Bool(&b, "val", "")

	// When
	err := y.Parse([]string{"-val=false"})

	// Then
	ts.Require().NoError(err)
	ts.Require().False(b)
}

type testFlagValue struct {
	called bool
}

func (tfv *testFlagValue) Set(val string) error {
	tfv.called = true
	return nil
}

func (tfv *testFlagValue) String() string {
	return "test_flag_value"
}

func (ts *TypesTestSuite) TestValue() {
	// Given
	val := &testFlagValue{}

	y := yag.New()
	y.Value(val, "val", "")

	// When
	err := y.Parse([]string{"-val=test_string"})

	// Then
	ts.Require().NoError(err)
	ts.Require().True(val.called)
	ts.Require().Equal("test_flag_value", val.String())
}
