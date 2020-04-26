package yag

import "flag"

type flagValue interface {
	flag.Value
	isSet() bool
}

type setterHelper struct {
	isSet bool
}

func (s *setterHelper) setStatus() {
	s.isSet = true
}

func (s *setterHelper) IsSet() bool {
	return s.isSet
}

func newStringValue(dest *string) *stringValue {
	return &stringValue{dest: dest}
}

type stringValue struct {
	setterHelper
	dest *string
}

func (s *stringValue) Set(val string) error {
	*s.dest = val
	s.setStatus()
	return nil
}

func (s *stringValue) String() string {
	if s.IsSet() {
		return *s.dest
	}
	return ""
}
