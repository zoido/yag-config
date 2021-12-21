package args

import (
	"github.com/zoido/yag-config/value"
)

// Strings returns new instance of Parser implementation that parses string values.
func Strings(s *[]string) Parser {
	return &stringsParser{dest: s}
}

type stringsParser struct {
	dest *[]string
}

func (sp *stringsParser) Parse(values []string) (int, error) {
	var n int
	outs := make([]string, len(values))
	for i, v := range values {
		n = i + 1
		o := value.String(&outs[i])
		// Cannot return error.
		_ = o.Set(v)
	}
	*sp.dest = outs
	return n, nil
}
