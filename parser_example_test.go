package yag_test

import (
	"fmt"

	"github.com/zoido/yag-config"
)

func ExampleParser_Usage() {
	type config struct {
		RequiredOption  string
		CustomEnvOption string
		NoFlagOption    string
		NoEnvOption     string
	}
	var arg int64
	var requiredArg int32
	var argWithName int16
	var requiredArgWithName int8
	var stringArgs []string

	y := yag.New(yag.WithEnvPrefix("MY_APP_"))
	cfg := &config{}

	y.String(&cfg.RequiredOption, "required_option", "sets required option", yag.Required())
	y.String(
		&cfg.CustomEnvOption,
		"custom_env_option",
		"sets custom env option",
		yag.FromEnv("MY_OPTION_TWO"),
	)
	y.String(&cfg.NoFlagOption, "no_flag_option", "sets np flag option", yag.NoFlag())
	y.String(&cfg.NoFlagOption, "no_env_option", "sets no env option", yag.NoEnv())
	y.Args().Int32(&requiredArg, yag.Required())
	y.Args().Int8(&requiredArgWithName, yag.WithName("my int8 argument"), yag.Required())
	y.Args().Int64(&arg)
	y.Args().Int16(&argWithName, yag.WithName("my int16 argument"))
	y.Args().Strings(&stringArgs)

	fmt.Print(y.Usage())

	// Output:
	//  int32 <my int8 argument> [int64] [my int16 argument] [string, ...]
	// 	-required_option ($MY_APP_REQUIRED_OPTION) [required]
	// 		sets required option
	// 	-custom_env_option ($MY_OPTION_TWO)
	// 		sets custom env option
	// 	$MY_APP_NO_FLAG_OPTION
	// 		sets np flag option
	// 	-no_env_option
	// 		sets no env option
}

func ExampleErrHelp() {
	var foo string

	y := yag.New()

	y.String(&foo, "foo", "sets Foo")

	err := y.Parse([]string{"--help"})
	if err == yag.ErrHelp {
		fmt.Printf("--help flag passed")
	}

	// Output: --help flag passed
}
