package yag_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zoido/yag-config"
)

func TestNew_Ok(t *testing.T) {
	// When
	yag.New()

	// Then
	// No panic.
}

func TestParse_Flags(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "flag value", str)
}

func TestParse_Env(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	t.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	require.NoError(t, err)
	require.Equal(t, "env value", str)
}

func TestParse_DefaultValue(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.Parse([]string{})

	// Then
	require.NoError(t, err)
	require.Equal(t, "default value", str)
}

func TestParse_WithEnvPrefix_Effective(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New(yag.WithEnvPrefix("MY_TEST_PREFIX_"))
	y.String(&str, "test_string", "sets test string value")
	t.Setenv("TEST_STRING", "env without prefix")
	t.Setenv("MY_TEST_PREFIX_TEST_STRING", "env with prefix")

	// When
	err := y.Parse([]string{})

	// Then
	require.NoError(t, err)
	require.Equal(t, "env with prefix", str)
}

func TestParse_FromEnv_Effective(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.FromEnv("DIFFERENT_TEST_ENV"))
	t.Setenv("TEST_STRING", "wrong env to look for")
	t.Setenv("DIFFERENT_TEST_ENV", "correct value")

	// When
	err := y.Parse([]string{})

	// Then
	require.NoError(t, err)
	require.Equal(t, "correct value", str)
}

func TestParseFlags(t *testing.T) {
	// Given
	var str1, str2 string

	y := yag.New()
	y.String(&str1, "test_string1", "sets test string value 1")
	y.String(&str1, "test_string2", "sets test string value 1")

	// When
	t.Setenv("TEST_STRING1", "env value 1")
	t.Setenv("TEST_STRING2", "env value 2")
	err := y.ParseFlags([]string{"-test_string1=flag value 1"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "flag value 1", str1)
	require.Equal(t, "", str2)
}

func TestParseEnv(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	t.Setenv("TEST_STRING", "env value")

	// When
	err := y.ParseEnv()

	// Then
	require.NoError(t, err)
	require.Equal(t, "env value", str)
}

func TestParseEnv_WithEnvPrefix_Effective(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New(yag.WithEnvPrefix("MY_TEST_PREFIX_"))
	y.String(&str, "test_string", "sets test string value")
	t.Setenv("TEST_STRING", "env without prefix")
	t.Setenv("MY_TEST_PREFIX_TEST_STRING", "env with prefix")

	// When
	err := y.ParseEnv()

	// Then
	require.NoError(t, err)
	require.Equal(t, "env with prefix", str)
}

func TestParse_FlagsTakePrecedence(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	t.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "flag value", str)
}

func TestParse_FlagsAlwaysTakePrecedence(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")
	t.Setenv("TEST_STRING", "env value")

	// When
	err := y.ParseFlags([]string{"-test_string=flag value"})
	require.NoError(t, err)
	err = y.ParseEnv()

	// Then
	require.NoError(t, err)
	require.Equal(t, "flag value", str)
}

func TestParse_RequiredOption_FailsOnParse(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.Parse([]string{})

	// Then
	require.Error(t, err)
	require.Contains(t, err.Error(), "required")
	require.Contains(t, err.Error(), "test_string")
}

func TestParse_RequiredOption_EnvEnough(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())
	t.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	require.NoError(t, err)
}

func TestParse_RequiredOption_FlagEnough(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	require.NoError(t, err)
}

func TestParseEnv_RequiredOption_FailsOnParse(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.ParseEnv()

	// Then
	require.Error(t, err)
}

func TestParseFlags_RequiredOption_FailsOnParse(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.Required())

	// When
	err := y.ParseFlags([]string{})

	// Then
	require.Error(t, err)
}

func TestParse_NoFlagOption_EnvOk(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.NoFlag())

	t.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	require.NoError(t, err)
	require.Equal(t, "env value", str)
}

func TestParse_NoFlagOption_FlagInvalid(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.NoFlag())

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	require.Error(t, err)
}

func TestParse_NoEnvOption_FlagOk(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.NoEnv())

	// When
	err := y.Parse([]string{"-test_string=flag value"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "flag value", str)
}

func TestParse_NoEnvOption_EnvIgnored(t *testing.T) {
	// Given
	str := "default value"

	y := yag.New()
	y.String(&str, "test_string", "sets test string value", yag.NoEnv())

	t.Setenv("TEST_STRING", "env value")

	// When
	err := y.Parse([]string{})

	// Then
	require.NoError(t, err)
	require.Equal(t, "default value", str)
}

func TestParseFlags_ErrorNotSwallowed(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.ParseFlags([]string{"--unknown_flag="})

	// Then
	require.Error(t, err)
}

func TestParseEnv_ErrorNotSwallowed(t *testing.T) {
	// Given
	var num int

	y := yag.New()
	y.Int(&num, "test_num", "sets test num value")
	t.Setenv("TEST_NUM", "invalid num value")

	// When
	err := y.ParseEnv()

	// Then
	require.Error(t, err)
}

func TestParse_Flags_ErrorNotSwallowed(t *testing.T) {
	// Given
	var str string

	y := yag.New()
	y.String(&str, "test_string", "sets test string value")

	// When
	err := y.Parse([]string{"--unknown_flag="})

	// Then
	require.Error(t, err)
}

func TestParse_Env_ErrorNotSwallowed(t *testing.T) {
	// Given
	var num int

	y := yag.New()
	y.Int(&num, "test_num", "sets test num value")
	t.Setenv("TEST_NUM", "invalid num value")

	// When
	err := y.Parse([]string{})

	// Then
	require.Error(t, err)
}
