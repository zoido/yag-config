package yag

import "flag"

type YagOption func(c *Yag)

func WithFlagSet(fs *flag.FlagSet) YagOption {
	return func(c *Yag) {
		c.flagSet = fs
	}
}

func WithEnvPrefix(prefix string) YagOption {
	return func(c *Yag) {
		c.envPrefix = prefix
	}
}

type VarOption func(v *variable)

func Required() VarOption {
	return func(v *variable) {
		v.required = true
	}
}

func FromEnv(envName string) VarOption {
	return func(v *variable) {
		v.envName = envName
	}
}
