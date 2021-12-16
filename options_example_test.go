package yag_test

import (
	"fmt"
	"os"

	"github.com/zoido/yag-config"
)

func ExampleWithEnvPrefix() {
	os.Clearenv()

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

func ExampleFromEnv() {
	os.Clearenv()

	type config struct {
		Foo string
	}

	y := yag.New()
	cfg := &config{}

	y.String(&cfg.Foo, "foo", "sets Foo", yag.FromEnv("DIFFERENT_FOO_VARIABLE"))

	_ = os.Setenv("DIFFERENT_FOO_VARIABLE", "Foo from different variable")
	_ = os.Setenv("FOO", "Foo from the default variable")

	err := y.Parse([]string{})
	if err != nil {
		os.Exit(2)
	}

	fmt.Printf(cfg.Foo)

	// Output: Foo from different variable
}

func ExampleRequired() {
	os.Clearenv()

	type config struct {
		Foo string
	}

	y := yag.New()
	cfg := &config{}

	y.String(&cfg.Foo, "foo", "sets Foo", yag.Required())

	err := y.Parse([]string{})

	fmt.Print(err)

	// Output: config option 'foo' is required
}

func ExampleRequired_flagOk() {
	os.Clearenv()

	type config struct {
		Foo string
	}

	y := yag.New()
	cfg := &config{}

	y.String(&cfg.Foo, "foo", "sets Foo", yag.Required())

	err := y.Parse([]string{"-foo=foo_value"})
	if err != nil {
		os.Exit(2)
	}

	fmt.Print(cfg.Foo)

	// Output: foo_value
}

func ExampleRequired_envOnlyOk() {
	os.Clearenv()

	type config struct {
		Foo string
	}

	y := yag.New()
	cfg := &config{}

	y.String(&cfg.Foo, "foo", "sets Foo", yag.Required())

	_ = os.Setenv("FOO", "foo_value")

	err := y.Parse([]string{})
	if err != nil {
		os.Exit(2)
	}

	fmt.Print(cfg.Foo)

	// Output: foo_value
}

func ExampleNoFlag() {
	os.Clearenv()

	type config struct {
		Foo string
	}

	y := yag.New()
	cfg := &config{}

	y.String(&cfg.Foo, "foo", "sets Foo", yag.NoFlag())

	err := y.Parse([]string{"-foo=foo_value"})

	fmt.Print(err)

	// Output: flag provided but not defined: -foo
}

func ExampleNoEnv() {
	os.Clearenv()

	type config struct {
		Foo string
	}

	y := yag.New()
	cfg := &config{
		Foo: "Default Foo value",
	}

	y.String(&cfg.Foo, "foo", "sets Foo", yag.NoEnv())

	_ = os.Setenv("FOO", "Foo from the environment variable ")
	err := y.Parse([]string{})
	if err != nil {
		os.Exit(2)
	}

	fmt.Printf(cfg.Foo)

	// Output: Default Foo value
}
