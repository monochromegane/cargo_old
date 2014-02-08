package group

import (
	"fmt"
	"github.com/monochromegane/cargo/cargo/docker"
	"github.com/monochromegane/cargo/cargo/option"
	"strings"
)

type GoPackage struct {
	From   string
	Option option.Option
}

func NewGoPackage(from string, option option.Option) *GoPackage {
	return &GoPackage{
		From:   from,
		Option: option,
	}
}

func (self *GoPackage) GroupBy() map[int][]string {
	opt := self.Option
	command := docker.RunCommand(
		opt.Image,
		docker.RunOption{
			SrcVolume: self.From,
			DstVolume: opt.Dest,
		},
		[]string{"/go/go/bin/go", "list", opt.GoPackage + "/..."},
	)
	result, err := command.Output()
	fmt.Printf(">>>> %s\n", result)
	if err != nil {
		return make(map[int][]string)
	}
	packages := strings.Split(string(result), "\n")

	return chunk(packages, opt.Concurrency)

}

func chunk(list []string, size int) map[int][]string {
	result := make(map[int][]string)
	loop := len(list) / size
	for i := 0; i < loop; i++ {
		max := (i * size) + size
		if max > len(list)-1 {
			max = len(list) - 1
		}
		result[i] = list[(i * size):max]
	}
	return result
}
