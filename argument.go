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
func (ap *ArgParser) Value(v flag.Value, options ...ArgOption) {
	ap.addArg(&wrapper{dest: v}, options...)
}

// String registers new string argument for parsing.
func (ap *ArgParser) String(s *string, options ...ArgOption) {
	ap.Value(value.String(s), options...)
}

// Strings tells parser to parse all of the leftover arguments as strings.
func (ap *ArgParser) Strings(s *[]string) {
	ap.parser = args.Strings(s)
}

func (ap *ArgParser) addArg(w *wrapper, options ...ArgOption) {
	a := &argument{
		value: w,
	}
	for _, opt := range options {
		opt.applyArg(a)
	}
	ap.args = append(ap.args, a)
}

// Parse parses the values according to the registered arguments.
func (ap *ArgParser) Parse(values []string) error {
	count := len(values)
	var nextToparse int
	for i, a := range ap.args {
		nextToparse = i + 1
		if nextToparse > count {
			break
		}
		v := values[i]

		err := a.value.Set(v)
		if err != nil {
			return fmt.Errorf("parsing argument '%s' on position %d", a.name, i+1)
		}
	}
	if (nextToparse < count) && (ap.parser != nil) {
		ap.parser.Parse(values[nextToparse:])
	}

	return ap.validate()
}

func (ap *ArgParser) validate() error {
	for i, a := range ap.args {
		if a.required && !a.value.isSet() {
			if a.name == "" {
				return fmt.Errorf("an argument is required on position %d", i+1)
			}
			return fmt.Errorf("argument '%s' is required on position %d", a.name, i+1)
		}
	}
	return nil
}

type argument struct {
	value *wrapper

	name     string
	required bool
}
