package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zoido/yag-config"
)

type config struct {
	Name *string
}

func main() {
	y := yag.NewYag()
	cfg := &config{
		Name: yag.String("nanana"),
	}

	os.Args = []string{"-name=aaa"}
	y.Add(&cfg.Name, "name", "sets the name")

	err := y.Parse(os.Args)
	if err != nil {
		log.Fatalf("err: %+v", err)
	}

	fmt.Printf(">> %+v << \n", *cfg.Name)
}
