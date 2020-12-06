# Yet Another Golang Config Library

<a href="https://pkg.go.dev/github.com/zoido/yag-config"><img src="https://pkg.go.dev/badge/github.com/zoido/yag-config"></a>
<a href="https://github.com/zoido/yag-config/blob/master/LICENSE"><img src="https://img.shields.io/github/license/zoido/yag-config?style=flat-square"/></a>
<img src="https://img.shields.io/github/workflow/status/zoido/yag-config/Go?style=flat-square&logoColor=white&logo=github"/>
<a href="https://codecov.io/gh/zoido/yag-config"><img src="https://img.shields.io/codecov/c/github/zoido/yag-config?style=flat-square&logoColor=white&logo=codecov"/></a>
<a href="https://goreportcard.com/report/github.com/zoido/yag-config"><img src="https://goreportcard.com/badge/github.com/zoido/yag-config?style=flat-square"></a>

## Overview

- obtain configuration values from flags and/or environment variables
  (flags always take precedence)
- any variable or struct member can become a destination for the config value
- define defaults in the native type
- option define environment variable prefix
- option to override the environment variable name
- option to mark options as required

## Example

<!-- markdownlint-disable MD010 -->

```go
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
y.Int(&cfg.Int, "int", "sets Int")

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
```

<!-- markdownlint-enable MD010 -->

## Supported types

- `str`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `bool`
- `time.Duration`
- any `flag.Value` implementation (e.g. [(github.com/sgreben/flagvar](https://github.com/sgreben/flagvar))
- more to come…

## Credits

Originally inspired by flags processing in of
[Consul agent](https://github.com/hashicorp/consul).
