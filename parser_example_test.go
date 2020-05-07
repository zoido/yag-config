package yag_test

import (
	"os"

	"github.com/zoido/yag-config"
)

func ExampleParser_Usage() {
	type config struct {
		Foo string
		Bar string
	}

	y := yag.New(yag.WithEnvPrefix("MY_APP_"))
	cfg := &config{
		Foo: "default foo value",
		Bar: "default bra value",
	}

	y.String(&cfg.Foo, "foo", "sets Foo")
	y.String(&cfg.Bar, "bar", "sets Bar", yag.FromEnv("MY_BAR"))

	err := y.Usage(os.Stdout)
	if err != nil {
		os.Exit(2)
	}

	// Output:
	// 	-foo ($MY_APP_FOO)
	// 		sets Foo
	// 	-bar ($MY_BAR)
	// 		sets Bar
}
