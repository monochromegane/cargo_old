package command

import (
	"os/exec"
)

type Command struct {
	Command string
	Args    []string
}

type Result struct {
	Output []byte
	Err    error
}

func (self *Command) Exec(out chan *Result) {
	cmd := exec.Command(self.Command, self.Args...)
	result, err := cmd.Output()
	out <- &Result{result, err}
}
