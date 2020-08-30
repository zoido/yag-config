package yag

import (
	"flag"

	"github.com/zoido/yag-config/value"
)

type flagValue interface {
	flag.Value
	isSet() bool
}

type flagWrapper struct {
	dest flag.Value
	b    bool
}

func (fw *flagWrapper) Set(val string) error {
	err := fw.dest.Set(val)
	if err != nil {
		return err
	}

	fw.b = true
	return nil
}

func (fw *flagWrapper) String() string {
	if fw.isSet() {
		return fw.dest.String()
	}
	return ""
}

func (fw *flagWrapper) IsBoolFlag() bool {
	return value.IsBoolFlag(fw.dest)
}

func (fw *flagWrapper) isSet() bool {
	return fw.b
}
