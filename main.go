package main

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/monochromegane/cargo/cargo"
	"github.com/monochromegane/cargo/cargo/option"
	"os"
)

var opts option.Option

func main() {

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "cargo"
	parser.Usage = "[OPTIONS]"

	_, err := parser.Parse()
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	cargo := cargo.Cargo{Option: opts}
	cargo.Run()

}
