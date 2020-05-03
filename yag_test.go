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
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("flag value", str)
}

func (s *YagTestSuite) TestParse_Env() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("env value", str)
}

func (s *YagTestSuite) TestParse_WithEnvPrefix_Effective() {
	// Given
	str := "default value"

	y := yag.New(yag.WithEnvPrefix("MY_TEST_PREFIX_"))
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env without prefix")
	os.Setenv("MY_TEST_PREFIX_TEST_STRING", "env with prefix")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("env with prefix", str)
}

func (s *YagTestSuite) TestParse_FromEnv_Effective() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.FromEnv("DIFFERENT_TEST_ENV"))
	os.Setenv("TEST_STRING", "wrong env to look for")
	os.Setenv("DIFFERENT_TEST_ENV", "correct value")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
	s.Require().Equal("correct value", str)
}

func (s *YagTestSuite) TestParseFlags() {
	// Given
	var str1, str2 string

	y := yag.New()
	y.String(&str1, "test_string1", "sets test string value 1")
	y.String(&str1, "test_string2", "sets test string value 1")

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
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.ParseEnv()

	// Then
	s.Require().NoError(err)
	s.Require().Equal("env value", str)
}

func (s *YagTestSuite) TestParseEnv_WithEnvPrefix_Effective() {
	// Given
	str := "default value"

	y := yag.New(yag.WithEnvPrefix("MY_TEST_PREFIX_"))
	y.String(&str, "test_string", "sets test string value")
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
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
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
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.ParseFlags([]string{"-test_string=flag value"})
	s.Require().NoError(err)
	err = y.ParseEnv()

	// Then
	s.Require().NoError(err)
	s.Require().Equal("flag value", str)
}

func (s *YagTestSuite) TestParse_Required_FailsOnParse() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "required")
	s.Require().Contains(err.Error(), "test_string")
}

func (s *YagTestSuite) TestParse_Required_EnvEnough() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	s.Require().NoError(err)
}

func (s *YagTestSuite) TestParse_Required_FlagEnough() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	s.Require().NoError(err)
}

func (s *YagTestSuite) TestParseEnv_Required_FailsOnParse() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.ParseEnv()

	// Then
	s.Require().Error(err)
}

func (s *YagTestSuite) TestParseFlags_Required_FailsOnParse() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.ParseFlags([]string{})

	// Then
	s.Require().Error(err)
}
