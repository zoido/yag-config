package args

// Parser defines the interface of commandline argument parser.
type Parser interface {
	Parse(values []string) error
}

// ParsingError is custom error returned by the parsers containing extra information about type
// and position where error ocurred
type ParsingError struct {
	Err error

	Position int
	Type     string
}

func (pe ParsingError) Error() string {
	return pe.Err.Error()
}
