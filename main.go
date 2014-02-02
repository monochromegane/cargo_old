package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"github.com/monochromegane/cargo/command"
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

	in := make(chan *command.Result)

	command := command.Command{"docker", []string{"run", opts.Image, opts.Command}}
	go command.Exec(in)

	for result := range in {
		fmt.Printf("%s, %s\n", result.Output, result.Err)
	}

}
