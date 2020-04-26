package yag_test

import (
	"fmt"
	"os"

	"github.com/zoido/yag-config"
)

type config struct {
	Foo string
	Bar string
	Baz string
	Qux string
}

func Example() {
	y := yag.New(yag.WithEnvPrefix("MY_APP_"))
	cfg := &config{
		Foo: "default foo value",
		Bar: "default bra value",
	}

	y.Register(&cfg.Foo, "foo", "sets Foo")
	y.Register(&cfg.Bar, "bar", "sets Bar")
	y.Register(&cfg.Baz, "baz", "sets Baz", yag.FromEnv("MY_BAZ_VALUE"))
	y.Register(&cfg.Qux, "qux", "sets Qux")

	args := []string{"-foo=foo flag value"}

	_ = os.Setenv("MY_APP_FOO", "foo env value")
	_ = os.Setenv("MY_APP_BAR", "bar env value")
	_ = os.Setenv("MY_BAZ_VALUE", "baz env value")

	err := y.Parse(args)
	if err != nil {
		os.Exit(2)
	}

	fmt.Printf("foo: %v\n", cfg.Foo)
	fmt.Printf("bar: %v\n", cfg.Bar)
	fmt.Printf("baz: %v\n", cfg.Baz)
	fmt.Printf("baz: %v\n", cfg.Qux)

	// Output:
	// foo: foo flag value
	// bar: bar env value
	// baz: baz env value
	// baz:
}
