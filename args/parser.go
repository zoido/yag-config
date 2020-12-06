package args

// Parser defines the interface of commandline argument parser.
type Parser interface {
	Parse(values []string) error
}
