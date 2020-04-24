package yag

// stringPtrValue is a flag.Value which stores the value in a *string.
// If the value was not set the pointer is nil.
type stringPtrValue struct {
	v **string
	b bool
}

func newStringPtrValue(p **string) *stringPtrValue {
	return &stringPtrValue{p, false}
}

func (s *stringPtrValue) Set(val string) error {
	*s.v, s.b = &val, true
	return nil
}

func (s *stringPtrValue) String() string {
	if s.b {
		return **s.v
	}
	return ""
}

func String(s string) *string {
	return &s
}
