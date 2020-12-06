package yag

// ParserOption configures the Parser.
type ParserOption func(c *Parser)

// WithEnvPrefix configures the prefix of the environment variable names that will be used for the
// lookup of the values. Default without prefix.
func WithEnvPrefix(prefix string) ParserOption {
	return func(c *Parser) {
		c.envPrefix = prefix
	}
}

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

// FromEnv overrides the environment variable name thad will be used to obtain the set value of the
// registered variable.
func FromEnv(envName string) VarOption {
	return &fromEnvOption{envName: envName}
}

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

type noEnvOption struct{}

func (*noEnvOption) applyVar(v *variable) {
	v.parseEnv = false
}

// NoFlag disables the flag for the variable. Useful for the options that should not appear on
// command line.
func NoFlag() VarOption {
	return &noFlagOption{}
}

type noFlagOption struct{}

func (*noFlagOption) applyVar(v *variable) {
	v.parseFlag = false
}
