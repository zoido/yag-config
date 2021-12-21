package yag_test

import (
	"fmt"
	"os"
	"time"

	"github.com/zoido/yag-config"
)

type config struct {
	Str      string
	Bool     bool
	Int      int
	Duration time.Duration
}

func Example() {
	y := yag.New(yag.WithEnvPrefix("MY_APP_"))
	cfg := &config{
		Str: "default str value",
		Int: 42,
	}

	y.String(&cfg.Str, "str", "sets Str")
	y.Bool(&cfg.Bool, "bool", "sets Bool")
	y.Duration(&cfg.Duration, "duration", "sets Duration", yag.FromEnv("MY_DURATION_VALUE"))
	y.Int(&cfg.Int, "int", "sets Int")

	var strArg string
	var intArgs []int
	y.Args().String(&strArg)
	y.Args().Ints(&intArgs)

	args := []string{"-str=str flag value", "str arg value", "3", "2", "1"}

	_ = os.Setenv("MY_APP_STR", "str env value")
	_ = os.Setenv("MY_APP_INT", "4")
	_ = os.Setenv("MY_DURATION_VALUE", "1h")

	err := y.Parse(args)
	if err != nil {
		os.Exit(2)
	}

	fmt.Printf("config.Str: %v\n", cfg.Str)
	fmt.Printf("config.Int: %v\n", cfg.Int)
	fmt.Printf("config.Bool: %v\n", cfg.Bool)
	fmt.Printf("config.Duration: %v\n", cfg.Duration)
	fmt.Printf("str arg: %v\n", strArg)
	fmt.Printf("int args: %v\n", intArgs)

	// Output:
	// config.Str: str flag value
	// config.Int: 4
	// config.Bool: false
	// config.Duration: 1h0m0s
	// str arg: str arg value
	// int args: [3 2 1]
}
