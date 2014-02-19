package group

import (
	"github.com/monochromegane/cargo/cargo/command"
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

	goList := command.GoListCommand{opt.Target}
	command := command.DockerRunCommand{
		Image:           opt.Image,
		HostVolume:      self.From,
		ContainerVolume: opt.Mount,
		WorkDir:         opt.Mount,
		Cmd:             goList.Command().Args,
	}
	result, err := command.Command().Output()
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
