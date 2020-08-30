package value_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zoido/yag-config/value"
)

func TestString(t *testing.T) {
	// Given
	var s string
	v := value.String(&s)

	// When
	err := v.Set("test string")

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal("test string", s)
	r.Equal("test string", v.String())
}
