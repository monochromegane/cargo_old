package concurrency

import (
        "fmt"
)

type Group interface {
	GroupBy([]string) map[int][]string
}

type Command interface {
	Output() ([]byte, error)
}

type Concurrency struct {
}

type CommandBuilder func(index int, args []string) Command

func (self *Concurrency) Run(group map[int][]string, builder CommandBuilder) {
        results := make(chan []byte)
        for index, args := range group {
                command := builder(index, args)
                go func(){
                        result, _ := command.Output()
                        results <- result
                }()
        }
        for i := 0; i < len(group); i++ {
                output := <-results
                fmt.Printf("%s\n", output)
        }
}
