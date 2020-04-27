package yag

import (
	"flag"
	"strconv"
)

// TODO:
//  [x] bool
//  [ ] string slice
//  [ ] Duration?

type flagValue interface {
	flag.Value
	isSet() bool
}

type isSetHelper struct {
	b bool
}

func (s *isSetHelper) setIsSet() {
	s.b = true
}

func (s *isSetHelper) isSet() bool {
	return s.b
}

type stringValue struct {
	isSetHelper
	dest *string
}

func (sv *stringValue) Set(val string) error {
	*sv.dest = val
	sv.setIsSet()
	return nil
}

func (sv *stringValue) String() string {
	if sv.isSet() {
		return *sv.dest
	}
	return ""
}

type intValue struct {
	isSetHelper
	dest *int
}

func (iv *intValue) Set(val string) error {
	num, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	*iv.dest = num
	iv.setIsSet()
	return nil
}

func (iv *intValue) String() string {
	if iv.isSet() {
		return strconv.Itoa(*iv.dest)
	}
	return ""
}

type boolValue struct {
	isSetHelper
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
	bv.setIsSet()
	return nil
}

func (bv *boolValue) String() string {
	if bv.isSet() {
		return strconv.FormatBool(*bv.dest)
	}
	return ""
}
