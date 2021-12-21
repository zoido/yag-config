package yag

// ParserOption configures the Parser.
type ParserOption func(c *Parser)

// VarOption configures handling of registered variables.
type VarOption interface {
	applyVar(v *variable)
}

// ArgOption configures handling of registered non-flag arguments.
type ArgOption interface {
	applyArg(v *argument)
}

// Option configures handling of registered variables or registered non-flag arguments.
type Option interface {
	VarOption
	ArgOption
}

// WithEnvPrefix configures the prefix of the environment variable names that will be used for the
// lookup of the values. Default without prefix.
func WithEnvPrefix(prefix string) ParserOption {
	return func(c *Parser) {
		c.envPrefix = prefix
	}
}

// FromEnv overrides the environment variable name thad will be used to obtain the set value of the
// registered variable.
func FromEnv(envName string) VarOption {
	return &fromEnvOption{envName: envName}
}

// FromEnv option implementation.
type fromEnvOption struct {
	envName string
}

func (feo *fromEnvOption) applyVar(v *variable) {
	v.envName = feo.envName
}

// Required sets the variable as required. Parsing will fail when the variable is not set via flags
// nor environment.
func Required() Option {
	return &requiredOption{}
}

type requiredOption struct{}

func (*requiredOption) applyVar(v *variable) {
	v.required = true
}

func (*requiredOption) applyArg(a *argument) {
	a.required = true
}

// NoEnv disables looking up of the variable value in the environment variables.
func NoEnv() VarOption {
	return &noEnvOption{}
}

// NoEnv option implementation.
type noEnvOption struct{}

func (*noEnvOption) applyVar(v *variable) {
	v.parseEnv = false
}

// NoFlag disables the flag for the variable. Useful for the options that should not appear on
// command line.
func NoFlag() VarOption {
	return &noFlagOption{}
}

// NoFlag option implementation.
type noFlagOption struct{}

func (*noFlagOption) applyVar(v *variable) {
	v.parseFlag = false
}

// WithName sets the name of the argument. This will be used in the usage message and errors.
func WithName(name string) ArgOption {
	return &withNameOption{name: name}
}

// Required option implementation.

// WithName option implementation.
type withNameOption struct {
	name string
}

func (wno *withNameOption) applyArg(a *argument) {
	a.name = wno.name
}
