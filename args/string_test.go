package args_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zoido/yag-config/args"
)

func TestString(t *testing.T) {
	// Given
	var s []string
	parser := args.Strings(&s)

	// When
	n, err := parser.Parse([]string{"a", "b", "c", "d"})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(4, n)
	r.Equal([]string{"a", "b", "c", "d"}, s)
}

func TestString_Empty(t *testing.T) {
	// Given
	s := []string{"x", "y", "z"}
	parser := args.Strings(&s)

	// When
	n, err := parser.Parse([]string{})

	// Then
	r := require.New(t)
	r.NoError(err)
	r.Equal(0, n)
	r.Equal([]string{}, s)
}
