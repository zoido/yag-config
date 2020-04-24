// Package yag is yet another configuration library for Go.
package yag

import (
	"flag"
	"fmt"
)

func NewYag(options ...Option) *Yag {
	y := &Yag{}
	for _, opt := range options {
		opt(y)
	}
	if y.flagSet == nil {
		y.flagSet = &flag.FlagSet{}
	}
	return y
}

type Yag struct {
	flagSet *flag.FlagSet

	err error
}

func (y *Yag) Add(p interface{}, name, help string) {
	if y.err != nil {
		return
	}
	switch x := p.(type) {
	case **string:
		y.flagSet.Var(newStringPtrValue(x), name, help)
	default:
		y.err = fmt.Errorf("invalid type: %T", p)
	}
}

func (y *Yag) Parse(args []string) error {
	if y.err != nil {
		return y.err
	}
	err := y.flagSet.Parse(args)
	if err != nil {
		return err
	}
	y.flagSet.PrintDefaults()
	return nil
}
