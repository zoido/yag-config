package args

// Parser defines the interface of commandline arguments parser.
type Parser interface {
	// Parse parses the values. Returns number of successfully parsed arguments and non-nil error
	// when parsing fails.
	Parse(values []string) (int, error)
}
