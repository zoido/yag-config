package yag

import (
	"flag"
	"fmt"

	"github.com/zoido/yag-config/args"
	"github.com/zoido/yag-config/value"
)

// ArgParser registers and parses non-flag arguments.
type ArgParser struct {
	args   []*argument
	parser args.Parser
}

// Value registers new generic flag.Value implementation for parsing an argument.
func (ap *ArgParser) Value(v flag.Value, name string, options ...ArgOption) {
	ap.addArg(&wrapper{dest: v}, name, options...)
}

// String registers new string argument for parsing.
func (ap *ArgParser) String(s *string, name string, options ...ArgOption) {
	ap.Value(value.String(s), name, options...)
}

// Strings TBD.
func (ap *ArgParser) Strings(s *[]string) {
	ap.parser = args.Strings(s)
}

func (ap *ArgParser) addArg(w *wrapper, name string, options ...ArgOption) {
	a := &argument{
		value: w,
		name:  name,
	}
	for _, opt := range options {
		opt.applyArg(a)
	}
	ap.args = append(ap.args, a)
}

func (ap *ArgParser) parse(values []string) error {
	count := len(values)
	var lastParsedIndex int
	for i, a := range ap.args {
		lastParsedIndex = i
		if i+1 > count {
			break
		}
		v := values[i]

		err := a.value.Set(v)
		if err != nil {
			return fmt.Errorf("parsing argument '%s' on position %d", a.name, i+1)
		}
	}
	if lastParsedIndex < count {
		ap.parser.Parse(values[lastParsedIndex:])
	}

	return ap.validate()
}

func (ap *ArgParser) validate() error {
	for _, a := range ap.args {
		if a.required && !a.value.isSet() {
			return fmt.Errorf("argument '%s' is required", a.name)
		}
	}
	return nil
}

type argument struct {
	value *wrapper

	name     string
	required bool
}
