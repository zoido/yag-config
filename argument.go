package yag

import (
	"flag"
	"fmt"
	"strings"

	"github.com/zoido/yag-config/args"
	"github.com/zoido/yag-config/value"
)

// ArgParser registers and parses non-flag arguments.
type ArgParser struct {
	args              []*argument
	parser            args.Parser
	parserPlaceholder string
}

// Value registers new generic flag.Value implementation for parsing an argument.
func (ap *ArgParser) Value(v flag.Value, options ...ArgOption) {
	ap.addArg(&wrapper{dest: v}, optionsWithPlaceholder("arg", options)...)
}

// String registers new string argument for parsing.
func (ap *ArgParser) String(s *string, options ...ArgOption) {
	ap.Value(value.String(s), optionsWithPlaceholder("string", options)...)
}

// Int registers new int argument for parsing.
func (ap *ArgParser) Int(i *int, options ...ArgOption) {
	ap.Value(value.Int(i), optionsWithPlaceholder("int", options)...)
}

// Int8 registers new int8 argument for parsing.
func (ap *ArgParser) Int8(i8 *int8, options ...ArgOption) {
	ap.Value(value.Int8(i8), optionsWithPlaceholder("int8", options)...)
}

// Int16 registers new int16 argument for parsing.
func (ap *ArgParser) Int16(i16 *int16, options ...ArgOption) {
	ap.Value(value.Int16(i16), optionsWithPlaceholder("int16", options)...)
}

// Int32 registers new int32 argument for parsing.
func (ap *ArgParser) Int32(i32 *int32, options ...ArgOption) {
	ap.Value(value.Int32(i32), optionsWithPlaceholder("int32", options)...)
}

// Int64 registers new int64 argument for parsing.
func (ap *ArgParser) Int64(i64 *int64, options ...ArgOption) {
	ap.Value(value.Int64(i64), optionsWithPlaceholder("int64", options)...)
}

// Strings tells parser to parse all of the leftover arguments as string.
func (ap *ArgParser) Strings(s *[]string) {
	ap.parser = args.Strings(s)
	ap.parserPlaceholder = "[string, ...]"
}

// Ints tells parser to parse all of the leftover arguments as int.
func (ap *ArgParser) Ints(i *[]int) {
	ap.parser = args.Ints(i)
	ap.parserPlaceholder = "[int, ...]"
}

// Int8s tells parser to parse all of the leftover arguments as int8.
func (ap *ArgParser) Int8s(i8 *[]int8) {
	ap.parser = args.Int8s(i8)
	ap.parserPlaceholder = "[int8, ...]"
}

// Int16s tells parser to parse all of the leftover arguments as int16.
func (ap *ArgParser) Int16s(i16 *[]int16) {
	ap.parser = args.Int16s(i16)
	ap.parserPlaceholder = "[int16, ...]"
}

// Int32s tells parser to parse all of the leftover arguments as int32.
func (ap *ArgParser) Int32s(i32 *[]int32) {
	ap.parser = args.Int32s(i32)
	ap.parserPlaceholder = "[int32, ...]"
}

// Int64s tells parser to parse all of the leftover arguments as int64.
func (ap *ArgParser) Int64s(i64 *[]int64) {
	ap.parser = args.Int64s(i64)
	ap.parserPlaceholder = "[int64, ...]"
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

func (ap *ArgParser) parse(values []string) error {
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
		if n, err := ap.parser.Parse(values[nextToparse:]); err != nil {
			return fmt.Errorf("parsing argument on position %d: %w", nextToparse+n+1, err)
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
			return fmt.Errorf("argument %q is required on position %d", a.name, i+1)
		}
	}
	return nil
}

func (ap *ArgParser) usage() string {
	u := make([]string, len(ap.args))
	for i, a := range ap.args {
		u[i] = a.usage()
	}
	u = append(u, ap.parserPlaceholder)
	return strings.Join(u, " ")
}

type argument struct {
	value *wrapper

	name        string
	placeholder string
	required    bool
}

func (a argument) usage() string {
	placeholder := a.placeholder
	if a.name != "" {
		placeholder = a.name
	}
	if a.name != "" && a.required {
		return "<" + a.name + ">"
	}
	if !a.required {
		placeholder = "[" + placeholder + "]"
	}
	return placeholder
}

func optionsWithPlaceholder(placeholder string, options []ArgOption) []ArgOption {
	return append([]ArgOption{&withPlaceholder{placeholder}}, options...)
}
