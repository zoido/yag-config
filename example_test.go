package yag_test

import (
	"fmt"
	"os"
	"time"

	"github.com/zoido/yag-config"
)

func Example() {
	type config struct {
		Str      string
		Bool     bool
		Int      int
		Duration time.Duration
	}

	y := yag.New(yag.WithEnvPrefix("MY_APP_"))
	cfg := &config{
		Str: "default str value",
		Int: 42,
	}

	y.String(&cfg.Str, "str", "sets Str")
	y.Bool(&cfg.Bool, "bool", "sets Bool")
	y.Duration(&cfg.Duration, "duration", "sets Duration", yag.FromEnv("MY_DURATION_VALUE"))
	y.Int(&cfg.Int, "int", "sets Qux")

	args := []string{"-str=str flag value"}

	_ = os.Setenv("MY_APP_STR", "str env value")
	_ = os.Setenv("MY_APP_INT", "4")
	_ = os.Setenv("MY_DURATION_VALUE", "1h")

	err := y.Parse(args)
	if err != nil {
		os.Exit(2)
	}

	fmt.Printf("config.Str: %v\n", cfg.Str)
	fmt.Printf("config.Int: %v\n", cfg.Int)
	fmt.Printf("config.Bool %v\n", cfg.Bool)
	fmt.Printf("config.Duration: %v\n", cfg.Duration)

	// Output:
	// config.Str: str flag value
	// config.Int: 4
	// config.Bool: false
	// config.Duration: 1h0m0s

	//
}

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
