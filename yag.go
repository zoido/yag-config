// Package yag is yet another configuration library for Go.
package yag

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func NewYag(options ...YagOption) *Yag {
	vars := make(map[string]*variable, 10)

	y := &Yag{
		vars: vars,
	}
	for _, opt := range options {
		opt(y)
	}
	if y.flagSet == nil {
		y.flagSet = &flag.FlagSet{}
	}
	return y
}

type Yag struct {
	envPrefix string
	flagSet   *flag.FlagSet

	err  error
	vars map[string]*variable
}

type variable struct {
	envName  string
	flagVal  flagValue
	help     string
	required bool
}

func (y *Yag) Add(p interface{}, name, help string, options ...VarOption) {
	if y.err != nil {
		return
	}
	switch x := p.(type) {
	case *string:
		y.addVar(newStringValue(x), name, help, options...)
	default:
		y.err = fmt.Errorf("unsupported type: %T; %w", p, y.err)
	}
}

func (y *Yag) addVar(val flagValue, name, help string, options ...VarOption) {
	variable := &variable{
		flagVal: val,
		envName: strings.ToUpper(fmt.Sprintf("%s%s", y.envPrefix, name)),
		help:    help,
	}
	for _, opt := range options {
		opt(variable)
	}
	y.vars[variable.envName] = variable
	y.flagSet.Var(val, name, help)
}

func (y *Yag) validate() error {
	return nil
}

func (y *Yag) ParseFlags(args []string) error {
	if err := y.parseFlags(args); err != nil {
		return err
	}
	return y.validate()
}

func (y *Yag) parseFlags(args []string) error {
	if y.err != nil {
		return y.err
	}
	return y.flagSet.Parse(args)
}

func (y *Yag) ParseEnv() error {
	if err := y.parseEnv(); err != nil {
		return err
	}
	return y.validate()

}

func (y *Yag) parseEnv() error {
	if y.err != nil {
		return y.err
	}

	for envName, variable := range y.vars {
		if value, envIsSet := os.LookupEnv(envName); envIsSet && !variable.flagVal.IsSet() {
			if err := variable.flagVal.Set(value); err != nil {
				return err
			}
		}
	}
	return nil
}

func (y *Yag) Parse(args []string) error {
	if y.err != nil {
		return y.err
	}
	if err := y.parseFlags(args); err != nil {
		return err
	}
	return y.parseEnv()
}
