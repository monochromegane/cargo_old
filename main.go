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
	parser.Usage = "[OPTIONS] COMMAND [ARG...]"

	args, err := parser.Parse()
	if err != nil || len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Printf("It works!\n")

	in := make(chan *command.Result)

	command := command.Command{"docker", append([]string{"run", opts.Image}, args...)}
	go command.Exec(in)

	for result := range in {
		fmt.Printf("%s, %s\n", result.Output, result.Err)
	}

}
