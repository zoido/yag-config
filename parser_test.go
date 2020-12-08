package yag_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/zoido/yag-config"
)

type ParserTestSuite struct {
	suite.Suite
}

func TestYag(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}

func (ts *ParserTestSuite) SetupTest() {
	os.Clearenv()
}

func (ts *ParserTestSuite) TestNew_Ok() {
	// When
	yag.New()

	// Then
	// No failure should occur.
}

func (ts *ParserTestSuite) TestParse_Flags() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("flag value", str)
}

func (ts *ParserTestSuite) TestParse_Env() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("env value", str)
}

func (ts *ParserTestSuite) TestParse_DefaultValue() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("default value", str)
}

func (ts *ParserTestSuite) TestParse_WithEnvPrefix_Effective() {
	// Given
	str := "default value"

	y := yag.New(yag.WithEnvPrefix("MY_TEST_PREFIX_"))
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env without prefix")
	os.Setenv("MY_TEST_PREFIX_TEST_STRING", "env with prefix")

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("env with prefix", str)
}

func (ts *ParserTestSuite) TestParse_FromEnv_Effective() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.FromEnv("DIFFERENT_TEST_ENV"))
	os.Setenv("TEST_STRING", "wrong env to look for")
	os.Setenv("DIFFERENT_TEST_ENV", "correct value")

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("correct value", str)
}

func (ts *ParserTestSuite) TestParseFlags() {
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
	ts.Require().NoError(err)
	ts.Require().Equal("flag value 1", str1)
	ts.Require().Equal("", str2)
}

func (ts *ParserTestSuite) TestParseEnv() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.ParseEnv()

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("env value", str)
}

func (ts *ParserTestSuite) TestParseEnv_WithEnvPrefix_Effective() {
	// Given
	str := "default value"

	y := yag.New(yag.WithEnvPrefix("MY_TEST_PREFIX_"))
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env without prefix")
	os.Setenv("MY_TEST_PREFIX_TEST_STRING", "env with prefix")

	// When
	err := y.ParseEnv()

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("env with prefix", str)
}

func (ts *ParserTestSuite) TestParse_FlagsTakePrecedence() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("flag value", str)
}

func (ts *ParserTestSuite) TestParse_FlagsAlwaysTakePrecedence() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.ParseFlags([]string{"-test_string=flag value"})
	ts.Require().NoError(err)
	err = y.ParseEnv()

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("flag value", str)
}

func (ts *ParserTestSuite) TestParse_RequiredOption_FailsOnParse() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().Error(err)
	ts.Require().Contains(err.Error(), "required")
	ts.Require().Contains(err.Error(), "test_string")
}

func (ts *ParserTestSuite) TestParse_RequiredOption_EnvEnough() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())
	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().NoError(err)
}

func (ts *ParserTestSuite) TestParse_RequiredOption_FlagEnough() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	ts.Require().NoError(err)
}

func (ts *ParserTestSuite) TestParseEnv_RequiredOption_FailsOnParse() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.ParseEnv()

	// Then
	ts.Require().Error(err)
}

func (ts *ParserTestSuite) TestParseFlags_RequiredOption_FailsOnParse() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.ParseFlags([]string{})

	// Then
	ts.Require().Error(err)
}

func (ts *ParserTestSuite) TestParse_NoFlagOption_EnvOk() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.NoFlag())

	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("env value", str)
}

func (ts *ParserTestSuite) TestParse_NoFlagOption_FlagInvalid() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.NoFlag())

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	ts.Require().Error(err)
}

func (ts *ParserTestSuite) TestParse_NoEnvOption_FlagOk() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.NoEnv())

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("flag value", str)
}

func (ts *ParserTestSuite) TestParse_NoEnvOption_EnvIgnored() {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.NoEnv())

	os.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().NoError(err)
	ts.Require().Equal("default value", str)
}

func (ts *ParserTestSuite) TestParseFlags_ErrorNotSwallowed() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.ParseFlags([]string{"--unknown_flag="})

	// Then
	ts.Require().Error(err)
}

func (ts *ParserTestSuite) TestParseEnv_ErrorNotSwallowed() {
	// Given
	var num int

	y := yag.New()
	y.Int(&num, "test_num", "sets test num value")
	os.Setenv("TEST_NUM", "invalid num value")

	// When
	err := y.ParseEnv()

	// Then
	ts.Require().Error(err)
}

func (ts *ParserTestSuite) TestParse_Flags_ErrorNotSwallowed() {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.Parse([]string{"--unknown_flag="})

	// Then
	ts.Require().Error(err)
}

func (ts *ParserTestSuite) TestParse_Env_ErrorNotSwallowed() {
	// Given
	var num int

	y := yag.New()
	y.Int(&num, "test_num", "sets test num value")
	os.Setenv("TEST_NUM", "invalid num value")

	// When
	err := y.Parse([]string{})

	// Then
	ts.Require().Error(err)
}
