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
	flagVal  flag.Value
	help     string
	required bool
}

func (y *Yag) Add(p interface{}, name, help string, options ...VarOption) {
	if y.err != nil {
		return
	}
	switch x := p.(type) {
	case **string:
		y.addVar(newStringPtrValue(x), name, help, options...)
	default:
		y.err = fmt.Errorf("unsupported type: %T; %w", p, y.err)
	}
}

func (y *Yag) addVar(flagVal flag.Value, name, help string, options ...VarOption) {
	variable := &variable{
		flagVal: flagVal,
		envName: strings.ToUpper(fmt.Sprintf("%s%s", y.envPrefix, name)),
		help:    help,
	}
	for _, opt := range options {
		opt(variable)
	}
	y.vars[variable.envName] = variable
	y.flagSet.Var(flagVal, name, help)
}

func (y *Yag) ParseFlags(args []string) error {
	if y.err != nil {
		return y.err
	}
	return y.flagSet.Parse(args)
}

func (y *Yag) ParseEnv() error {
	if y.err != nil {
		return y.err
	}

	for envName, variable := range y.vars {
		if value, isSet := os.LookupEnv(envName); isSet {
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
	if err := y.ParseEnv(); err != nil {
		return err
	}
	return y.ParseFlags(args)
}
