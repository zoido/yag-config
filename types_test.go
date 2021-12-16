package yag_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zoido/yag-config"
)

func TestString(t *testing.T) {
	// Given
	var s string

	y := yag.New()
	y.String(&s, "val", "")

	// When
	err := y.Parse([]string{"-val=test_string"})

	// Then
	require.NoError(t, err)
	require.Equal(t, "test_string", s)
}

func TestInt(t *testing.T) {
	// Given
	var n int

	y := yag.New()
	y.Int(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, 8, n)
}

func TestInt8(t *testing.T) {
	// Given
	var n int8

	y := yag.New()
	y.Int8(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, int8(8), n)
}

func TestInt16(t *testing.T) {
	// Given
	var n int16

	y := yag.New()
	y.Int16(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, int16(8), n)
}

func TestInt32(t *testing.T) {
	// Given
	var n int32

	y := yag.New()
	y.Int32(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, int32(8), n)
}

func TestInt64(t *testing.T) {
	// Given
	var n int64

	y := yag.New()
	y.Int64(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, int64(8), n)
}

func TestUint(t *testing.T) {
	// Given
	var n uint

	y := yag.New()
	y.Uint(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, uint(8), n)
}

func TestUint8(t *testing.T) {
	// Given
	var n uint8

	y := yag.New()
	y.Uint8(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, uint8(8), n)
}

func TestUint16(t *testing.T) {
	// Given
	var n uint16

	y := yag.New()
	y.Uint16(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, uint16(8), n)
}

func TestUint32(t *testing.T) {
	// Given
	var n uint32

	y := yag.New()
	y.Uint32(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, uint32(8), n)
}

func TestUint64(t *testing.T) {
	// Given
	var n uint64

	y := yag.New()
	y.Uint64(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=8"})

	// Then
	require.NoError(t, err)
	require.Equal(t, uint64(8), n)
}

func TestFloat32(t *testing.T) {
	// Given
	var n float32

	y := yag.New()
	y.Float32(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=6.626E-34"})

	// Then
	require.NoError(t, err)
	require.Equal(t, float32(6.626e-34), n)
}

func TestFloat64(t *testing.T) {
	// Given
	var n float64

	y := yag.New()
	y.Float64(&n, "val", "")

	// When
	err := y.Parse([]string{"-val=6.626E-34"})

	// Then
	require.NoError(t, err)
	require.Equal(t, float64(6.626e-34), n)
}

func TestDuration(t *testing.T) {
	// Given
	var d time.Duration

	y := yag.New()
	y.Duration(&d, "val", "")

	// When
	err := y.Parse([]string{"-val=10h30m15s"})

	// Then
	require.NoError(t, err)
	require.Equal(t, float64(10*60*60+30*60+15), d.Seconds())
}

func TestBool(t *testing.T) {
	// Given
	var b bool

	y := yag.New()
	y.Bool(&b, "val", "")

	// When
	err := y.Parse([]string{"-val=true"})

	// Then
	require.NoError(t, err)
	require.True(t, b)
}

func TestBool_BoolFlag(t *testing.T) {
	// Given
	var b bool

	y := yag.New()
	y.Bool(&b, "val", "")

	// When
	err := y.Parse([]string{"-val"})

	// Then
	require.NoError(t, err)
	require.True(t, b)
}

func TestBool_False(t *testing.T) {
	// Given
	b := false

	y := yag.New()
	y.Bool(&b, "val", "")

	// When
	err := y.Parse([]string{"-val=false"})

	// Then
	require.NoError(t, err)
	require.False(t, b)
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

func TestValue(t *testing.T) {
	// Given
	val := &testFlagValue{}

	y := yag.New()
	y.Value(val, "val", "")

	// When
	err := y.Parse([]string{"-val=test_string"})

	// Then
	require.NoError(t, err)
	require.True(t, val.called)
	require.Equal(t, "test_flag_value", val.String())
}
