// Package yag is yet another configuration library for Go.
package yag

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Parser registers and parses configuration values.
type Parser struct {
	envPrefix string
	flagSet   *flag.FlagSet

	vars []*value
}

// New returns new instance of the Yag.
func New(options ...ParserOption) *Parser {
	vars := make([]*value, 0, 10)

	y := &Parser{
		vars: vars,
	}
	for _, opt := range options {
		opt(y)
	}
	y.flagSet = &flag.FlagSet{Usage: func() {}}
	return y
}

type value struct {
	name      string
	envName   string
	flagVal   flagValue
	help      string
	required  bool
	parseFlag bool
	parseEnv  bool
}

func (v *value) usage() string {
	u := []string{"\t"}
	if v.parseFlag {
		u = append(u, "-", v.name)
	}
	if v.parseEnv && v.parseFlag {
		u = append(u, " ($", v.envName, ")")
	}
	if v.parseEnv && !v.parseFlag {
		u = append(u, "$", v.envName)
	}
	if v.required {
		u = append(u, " [required]")
	}
	u = append(u, "\n\t\t", v.help)

	return strings.Join(u, "")
}

// Value registers new generic flag.Value implementation for parsing.
func (y *Parser) Value(v flag.Value, name, help string, options ...VarOption) {
	y.addVar(&flagWrapper{dest: v}, name, help, options...)
}

// String registers new string variable for parsing.
func (y *Parser) String(s *string, name, help string, options ...VarOption) {
	y.Value(&stringValue{dest: s}, name, help, options...)
}

// Int registers new int variable for parsing.
func (y *Parser) Int(i *int, name, help string, options ...VarOption) {
	y.Value(&intValue{dest: i}, name, help, options...)
}

// Int8 registers new int8 variable for parsing.
func (y *Parser) Int8(i *int8, name, help string, options ...VarOption) {
	y.Value(&int8Value{dest: i}, name, help, options...)
}

// Int16 registers new int16 variable for parsing.
func (y *Parser) Int16(i *int16, name, help string, options ...VarOption) {
	y.Value(&int16Value{dest: i}, name, help, options...)
}

// Int32 registers new int32 variable for parsing.
func (y *Parser) Int32(i *int32, name, help string, options ...VarOption) {
	y.Value(&int32Value{dest: i}, name, help, options...)
}

// Int64 registers new int64 variable for parsing.
func (y *Parser) Int64(i *int64, name, help string, options ...VarOption) {
	y.Value(&int64Value{dest: i}, name, help, options...)
}

// Uint registers new uint variable for parsing.
func (y *Parser) Uint(i *uint, name, help string, options ...VarOption) {
	y.Value(&uintValue{dest: i}, name, help, options...)
}

// Uint8 registers new uint8 variable for parsing.
func (y *Parser) Uint8(i *uint8, name, help string, options ...VarOption) {
	y.Value(&uint8Value{dest: i}, name, help, options...)
}

// Uint16 registers new uint16 variable for parsing.
func (y *Parser) Uint16(i *uint16, name, help string, options ...VarOption) {
	y.Value(&uint16Value{dest: i}, name, help, options...)
}

// Uint32 registers new uint32 variable for parsing.
func (y *Parser) Uint32(i *uint32, name, help string, options ...VarOption) {
	y.Value(&uint32Value{dest: i}, name, help, options...)
}

// Uint64 registers new uint64 variable for parsing.
func (y *Parser) Uint64(i *uint64, name, help string, options ...VarOption) {
	y.Value(&uint64Value{dest: i}, name, help, options...)
}

// Float32 registers new float32 variable for parsing.
func (y *Parser) Float32(i *float32, name, help string, options ...VarOption) {
	y.Value(&float32Value{dest: i}, name, help, options...)
}

// Float64 registers new float64 variable for parsing.
func (y *Parser) Float64(i *float64, name, help string, options ...VarOption) {
	y.Value(&float64Value{dest: i}, name, help, options...)
}

// Bool registers new bool variable for parsing.
func (y *Parser) Bool(b *bool, name, help string, options ...VarOption) {
	y.Value(&boolValue{dest: b}, name, help, options...)
}

// Duration registers new time.Duration variable for parsing.
func (y *Parser) Duration(d *time.Duration, name, help string, options ...VarOption) {
	y.Value(&durationValue{dest: d}, name, help, options...)
}

func (y *Parser) addVar(val flagValue, name, help string, options ...VarOption) {
	variable := &value{
		flagVal:   val,
		envName:   strings.ToUpper(fmt.Sprintf("%s%s", y.envPrefix, name)),
		name:      name,
		help:      help,
		parseEnv:  true,
		parseFlag: true,
	}
	for _, opt := range options {
		opt(variable)
	}
	y.vars = append(y.vars, variable)
	if variable.parseFlag {
		y.flagSet.Var(val, name, help)
	}
}

func (y *Parser) validate() error {
	for _, variable := range y.vars {
		if variable.required && !variable.flagVal.isSet() {
			return fmt.Errorf("config option '%s' is required", variable.name)
		}
	}
	return nil
}

// ParseFlags parses configuration values from commandline flags.
func (y *Parser) ParseFlags(args []string) error {
	if err := y.doParseFlags(args); err != nil {
		return err
	}
	return y.validate()
}

func (y *Parser) doParseFlags(args []string) error {
	return y.flagSet.Parse(args)
}

// ParseEnv parses configuration values from environment variables.
func (y *Parser) ParseEnv() error {
	if err := y.doParseEnv(); err != nil {
		return err
	}
	return y.validate()
}

func (y *Parser) doParseEnv() error {
	for _, v := range y.vars {
		if !v.parseEnv {
			continue
		}

		value, envIsSet := os.LookupEnv(v.envName)
		if envIsSet && !v.flagVal.isSet() {
			if err := v.flagVal.Set(value); err != nil {
				return err
			}
		}
	}
	return nil
}

// Parse parses configuration values from flags and environment variables.
// Flags values always override environment variable value.
func (y *Parser) Parse(args []string) error {
	if err := y.doParseFlags(args); err != nil {
		return err
	}
	if err := y.doParseEnv(); err != nil {
		return err
	}
	return y.validate()
}

// Usage returns formatted string with usage help.
func (y *Parser) Usage() string {
	u := make([]string, 0, len(y.vars))
	for _, v := range y.vars {
		u = append(u, v.usage())
	}
	return strings.Join(u, "\n")
}
