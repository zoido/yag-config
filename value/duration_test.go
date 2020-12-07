package value_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zoido/yag-config/value"
)

func TestDuration(t *testing.T) {
	// Given
	var d time.Duration
	v := value.Duration(&d)

	// When
	err := v.Set("10h")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(float64(10*60*60), d.Seconds())
	r.Equal(d.String(), v.String())
}

func TestDuration_ParseError(t *testing.T) {
	// Given
	var d time.Duration
	v := value.Duration(&d)

	// When
	err := v.Set("10x")

	// Then
	r := require.New(t)
	r.Error(err)
	r.Contains(err.Error(), "10x")
}
