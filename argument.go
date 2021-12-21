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

// Int registers new int argument for parsing.
func (ap *ArgParser) Int(i *int, options ...ArgOption) {
	ap.Value(value.Int(i), options...)
}

// Int8 registers new int8 argument for parsing.
func (ap *ArgParser) Int8(i8 *int8, options ...ArgOption) {
	ap.Value(value.Int8(i8), options...)
}

// Int16 registers new int16 argument for parsing.
func (ap *ArgParser) Int16(i16 *int16, options ...ArgOption) {
	ap.Value(value.Int16(i16), options...)
}

// Int32 registers new int32 argument for parsing.
func (ap *ArgParser) Int32(i32 *int32, options ...ArgOption) {
	ap.Value(value.Int32(i32), options...)
}

// Int64 registers new int64 argument for parsing.
func (ap *ArgParser) Int64(i64 *int64, options ...ArgOption) {
	ap.Value(value.Int64(i64), options...)
}

// Strings tells parser to parse all of the leftover arguments as string.
func (ap *ArgParser) Strings(s *[]string) {
	ap.parser = args.Strings(s)
}

// Ints tells parser to parse all of the leftover arguments as int.
func (ap *ArgParser) Ints(i *[]int) {
	ap.parser = args.Ints(i)
}

// Int8s tells parser to parse all of the leftover arguments as int8.
func (ap *ArgParser) Int8s(i8 *[]int8) {
	ap.parser = args.Int8s(i8)
}

// Int16s tells parser to parse all of the leftover arguments as int16.
func (ap *ArgParser) Int16s(i16 *[]int16) {
	ap.parser = args.Int16s(i16)
}

// Int32s tells parser to parse all of the leftover arguments as int32.
func (ap *ArgParser) Int32s(i32 *[]int32) {
	ap.parser = args.Int32s(i32)
}

// Int64s tells parser to parse all of the leftover arguments as int64.
func (ap *ArgParser) Int64s(i64 *[]int64) {
	ap.parser = args.Int64s(i64)
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

		if err := a.value.Set(v); err != nil {
			if a.name == "" {
				return fmt.Errorf("parsing argument on position %d: %w", i+1, err)
			}
			return fmt.Errorf("parsing argument '%s' on position %d: %w", a.name, i+1, err)
		}
	}
	if (nextToparse < count) && (ap.parser != nil) {
		if err := ap.parser.Parse(values[nextToparse:]); err != nil {
			if pErr, ok := err.(args.ParsingError); ok {
				return fmt.Errorf(
					"parsing %s argument on position %d: %w",
					pErr.Type,
					pErr.Position+nextToparse,
					pErr.Err,
				)
			}
			return err
		}
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
