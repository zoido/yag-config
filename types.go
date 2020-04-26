package yag

import "flag"

type flagValue interface {
	flag.Value
	IsSet() bool
}

func newStringValue(dest *string) *stringValue {
	return &stringValue{dest: dest}
}

type stringValue struct {
	dest  *string
	isSet bool
}

func (s *stringValue) Set(val string) error {
	*s.dest, s.isSet = val, true
	return nil
}

func (s *stringValue) String() string {
	if s.isSet {
		return *s.dest
	}
	return ""
}

func (s *stringValue) IsSet() bool {
	return s.isSet
}
