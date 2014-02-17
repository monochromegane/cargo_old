package concurrency

type Commander interface {
	Output() ([]byte, error)
}

type Result struct {
	Output []byte
	Err    error
}

type CommandBuilder func(index int, args []string) Commander
type ResultFunc func(index int, args []string, result []byte, err error)

func Run(group map[int][]string, builder CommandBuilder, result ResultFunc) {
	results := make(chan Result)
	for index, args := range group {
		command := builder(index, args)
		go func(command Commander, results chan Result) {
			result, err := command.Output()
			results <- Result{result, err}
		}(command, results)
	}
	for i := 0; i < len(group); i++ {
		r := <-results
		result(i, group[i], r.Output, r.Err)
	}
}
