package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"github.com/monochromegane/cargo/option"
	"os"
)

var opts option.Option

func main() {

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "cargo"

	_, err := parser.Parse()
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Printf("It works!\n")

}
