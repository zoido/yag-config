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

	err  error
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
	name     string
	envName  string
	flagVal  flagValue
	help     string
	required bool
}

// Register registers new variable for parsing.
func (y *Parser) Register(ref interface{}, name, help string, options ...VarOption) {
	switch x := ref.(type) {
	case *string:
		y.addVar(&stringValue{dest: x}, name, help, options...)
	case *int:
		y.addVar(&intValue{dest: x}, name, help, options...)
	case *bool:
		y.addVar(&boolValue{dest: x}, name, help, options...)
	case *time.Duration:
		y.addVar(&durationValue{dest: x}, name, help, options...)
	default:
		msg := fmt.Sprintf("unsupported type: %s(%T)", name, ref)
		if y.err == nil {
			y.err = fmt.Errorf(msg)
		} else {
			y.err = fmt.Errorf("%s, %s", y.err.Error(), msg)
		}
	}
}

func (y *Parser) addVar(val flagValue, name, help string, options ...VarOption) {
	variable := &value{
		flagVal: val,
		envName: strings.ToUpper(fmt.Sprintf("%s%s", y.envPrefix, name)),
		name:    name,
		help:    help,
	}
	for _, opt := range options {
		opt(variable)
	}
	y.vars = append(y.vars, variable)
	y.flagSet.Var(val, name, help)
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
	if y.err != nil {
		return y.err
	}
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
	if y.err != nil {
		return y.err
	}

	for _, v := range y.vars {
		if value, envIsSet := os.LookupEnv(v.envName); envIsSet && !v.flagVal.isSet() {
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
	if y.err != nil {
		return y.err
	}
	if err := y.doParseFlags(args); err != nil {
		return err
	}
	if err := y.doParseEnv(); err != nil {
		return err
	}
	return y.validate()
}

// Usage outputs flag and environment value usage to the file.
func (y *Parser) Usage(f *os.File) error {
	for _, v := range y.vars {
		_, err := fmt.Fprintf(f, "\t-%s ($%s)\n\t\t%s\n", v.name, v.envName, v.help)
		if err != nil {
			return err
		}
	}
	return nil
}
