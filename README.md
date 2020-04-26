# Yet Another Golang Config Library

![Go](https://github.com/zoido/yag-config/workflows/Go/badge.svg)

## Overview

- obtain configuration values from flags and/or environment variables
  (flags always take precedence)
- any variable or struct member can become a destination for the config value
- define defaults in the native type
- option define environment variable prefix
- option to override the environment variable name

## Example

<!-- markdownlint-disable MD010 -->

```go
type config struct {
	Foo string
	Bar string
	Baz string
	Qux string
}

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

os.Setenv("MY_APP_FOO", "foo env value")
os.Setenv("MY_APP_BAR", "bar env value")
os.Setenv("MY_BAZ_VALUE", "baz env value")

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
```

<!-- markdownlint-enable MD010 -->

## Supported types

- `str`
- `int`
- more to come…
