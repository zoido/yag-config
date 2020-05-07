package yag_test

import (
	"fmt"
	"os"

	"github.com/zoido/yag-config"
)

func Example_WithEnvPrefix() {
	type config struct {
		Foo string
	}

	y := yag.New(yag.WithEnvPrefix("MY_APP_"))
	cfg := &config{}

	y.String(&cfg.Foo, "foo", "sets Foo")

	_ = os.Setenv("MY_APP_FOO", "Foo from the variable with prefix")
	_ = os.Setenv("FOO", "Foo from the variable without prefix")

	err := y.Parse([]string{})

	if err != nil {
		os.Exit(2)
	}

	fmt.Printf(cfg.Foo)

	// Output: Foo from the variable with prefix
}
