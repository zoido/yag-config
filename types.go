package yag

import "flag"

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
