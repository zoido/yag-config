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

	args := []string{"-str=str flag value"}

	_ = os.Setenv("MY_APP_STR", "str env value")
	_ = os.Setenv("MY_APP_INT", "4")
	_ = os.Setenv("MY_DURATION_VALUE", "1h")

	err := y.Parse(args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("config.Str: %v, ", cfg.Str)
	fmt.Printf("config.Int: %v, ", cfg.Int)
	fmt.Printf("config.Bool: %v, ", cfg.Bool)
	fmt.Printf("config.Duration: %v", cfg.Duration)

	// Output:
	// config.Str: str flag value, config.Int: 4, config.Bool: false, config.Duration: 1h0m0s
}
