package yag

import (
	"flag"
	"strconv"
)

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

func newStringValue(dest *string) *stringValue {
	return &stringValue{dest: dest}
}

type stringValue struct {
	isSetHelper
	dest *string
}

func (s *stringValue) Set(val string) error {
	*s.dest = val
	s.setIsSet()
	return nil
}

func (s *stringValue) String() string {
	if s.isSet() {
		return *s.dest
	}
	return ""
}

func newIntValue(dest *int) *intValue {
	return &intValue{dest: dest}
}

type intValue struct {
	isSetHelper
	dest *int
}

func (s *intValue) Set(val string) error {
	num, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	*s.dest = num
	s.setIsSet()
	return nil
}

func (s *intValue) String() string {
	if s.isSet() {
		return strconv.Itoa(*s.dest)
	}
	return ""
}
