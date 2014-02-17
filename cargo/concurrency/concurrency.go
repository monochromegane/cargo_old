package concurrency

import (
	"os/exec"
)

type Result struct {
	Output []byte
	Err    error
}

type RunCommand func(index int, args []string) *exec.Cmd
type OnResult func(index int, args []string, result []byte, err error) bool

func Run(group map[int][]string, run RunCommand, onResult OnResult) {
	results := make(chan Result)
	for index, args := range group {
		command := run(index, args)
		go func(command *exec.Cmd, results chan Result) {
			result, err := command.Output()
			results <- Result{result, err}
		}(command, results)
	}
	for i := 0; i < len(group); i++ {
		r := <-results
		if !onResult(i, group[i], r.Output, r.Err) {
			break
		}
	}
}
