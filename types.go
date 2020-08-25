package yag

import (
	"flag"
	"strconv"
	"time"
)

type flagValue interface {
	flag.Value
	isSet() bool
}

// optional interface to indicate boolean flags that can be
// supplied without "=value" text
type boolFlag interface {
	flag.Value
	IsBoolFlag() bool
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
	if fv, ok := fw.dest.(boolFlag); ok {
		return fv.IsBoolFlag()
	}
	return false
}

func (fw *flagWrapper) isSet() bool {
	return fw.b
}

type stringValue struct {
	dest *string
}

func (sv *stringValue) Set(val string) error {
	*sv.dest = val
	return nil
}

func (sv *stringValue) String() string {
	return *sv.dest
}

type intValue struct {
	dest *int
}

func (iv *intValue) Set(val string) error {
	num, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	*iv.dest = num
	return nil
}

func (iv *intValue) String() string {
	return strconv.Itoa(*iv.dest)
}

type int64Value struct {
	dest *int64
}

func (iv *int64Value) Set(val string) error {
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}

	*iv.dest = num
	return nil
}

func (iv *int64Value) String() string {
	return strconv.FormatInt(*iv.dest, 10)
}

type boolValue struct {
	dest *bool
}

func (*boolValue) IsBoolFlag() bool {
	return true
}

func (bv *boolValue) Set(val string) error {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}

	*bv.dest = b
	return nil
}

func (bv *boolValue) String() string {
	return strconv.FormatBool(*bv.dest)
}

type durationValue struct {
	dest *time.Duration
}

func (dv *durationValue) Set(val string) error {
	duration, err := time.ParseDuration(val)
	if err != nil {
		return err
	}

	*dv.dest = duration
	return nil
}

func (dv *durationValue) String() string {
	return dv.dest.String()
}
