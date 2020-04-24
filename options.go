package yag

import "flag"

type Option func(c *Yag)

func WithFlagSet(fs *flag.FlagSet) Option {
	return func(c *Yag) {
		c.flagSet = fs
	}
}
