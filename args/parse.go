package args

import (
	"fmt"

	"github.com/zoido/yag-config/value"
)

type Parser interface {
	Parse(values []string) error
}

func Strings(s *[]string) Parser {
	return &stringsParser{dest: s}
}

type stringsParser struct {
	dest *[]string
}

func (sp *stringsParser) Parse(values []string) error {
	outs := make([]string, len(values))
	for i, v := range values {
		o := value.String(&outs[i])
		err := o.Set(v)
		if err != nil {
			return fmt.Errorf("parsing argument on position %d", i+1)
		}
	}
	sp.dest = &outs
	return nil
}
