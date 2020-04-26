package yag_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/zoido/yag-config"
)

type YagTestSuite struct {
	suite.Suite
}

func TestYagTestSuite(t *testing.T) {
	suite.Run(t, new(YagTestSuite))
}

func (s *YagTestSuite) SetupTest() {
	os.Clearenv()
}

func (s *YagTestSuite) TestNew_Ok() {
	// When
	yag.New()

	// Then
	// No failure should occur.
}

func (s *YagTestSuite) TestParse_Flags() {
	// Given
	var str string

	y := yag.New()
	y.Register(&str, "test_string", "sets test string value")

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("flag value", str)
}

func (s *YagTestSuite) TestParse_Env() {
	// Given
	var str string

	y := yag.New()
	y.Register(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("env value", str)
}

func (s *YagTestSuite) TestParse_WithEnvPrefixEffective() {
	// Given
	var str string

	y := yag.New(yag.WithEnvPrefix("MY_TEST_PREFIX_"))
	y.Register(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env without prefix")
	os.Setenv("MY_TEST_PREFIX_TEST_STRING", "env with prefix")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("env with prefix", str)
}

func (s *YagTestSuite) TestParseFlags() {
	// Given
	var str1, str2 string

	y := yag.New()
	y.Register(&str1, "test_string1", "sets test string value 1")
	y.Register(&str1, "test_string2", "sets test string value 1")

	// When
	os.Setenv("TEST_STRING1", "env value 1")
	os.Setenv("TEST_STRING2", "env value 2")
	err := y.ParseFlags([]string{"-test_string1=flag value 1"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("flag value 1", str1)
	s.Require().Equal("", str2)
}

func (s *YagTestSuite) TestParseEnv() {
	// Given
	var str string

	y := yag.New()
	y.Register(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.ParseEnv()

	// Then
	s.Require().NoError(err)
	s.Require().Equal("env value", str)
}

func (s *YagTestSuite) TestParseEnv_WithEnvPrefixEffective() {
	// Given
	var str string

	y := yag.New(yag.WithEnvPrefix("MY_TEST_PREFIX_"))
	y.Register(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env without prefix")
	os.Setenv("MY_TEST_PREFIX_TEST_STRING", "env with prefix")

	// When
	err := y.ParseEnv()

	// Then
	s.Require().NoError(err)
	s.Require().Equal("env with prefix", str)
}

func (s *YagTestSuite) TestParse_FlagsTakePrecedence() {
	// Given
	var str string

	y := yag.New()
	y.Register(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("flag value", str)
}

func (s *YagTestSuite) TestParse_FlagsAlwaysTakePrecedence() {
	// Given
	var str string

	y := yag.New()
	y.Register(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.ParseFlags([]string{"-test_string=flag value"})
	s.Require().NoError(err)
	err = y.ParseEnv()

	// Then
	s.Require().NoError(err)
	s.Require().Equal("flag value", str)
}

func (s *YagTestSuite) TestRegister_UnsupportedType_ParseError() {
	// Given
	type testType struct{}
	testVar := &testType{}

	y := yag.New()
	y.Register(testVar, "test_type", "sets test type var")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "unsupported type")
	s.Require().Contains(err.Error(), "test_type")
	s.Require().Contains(err.Error(), "testType")
}

func (s *YagTestSuite) TestRegister_UnsupportedType_ParseError_AllInTheError() {
	// Given
	type testType1 struct{}
	type testType2 struct{}
	testVar1 := &testType1{}
	testVar2 := &testType2{}

	y := yag.New()
	y.Register(testVar1, "test_type1", "sets test type 1 var")
	y.Register(testVar2, "test_type2", "sets test type 2 var")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "unsupported type")

	s.Require().Contains(err.Error(), "test_type1")
	s.Require().Contains(err.Error(), "testType1")

	s.Require().Contains(err.Error(), "test_type2")
	s.Require().Contains(err.Error(), "testType2")
}
