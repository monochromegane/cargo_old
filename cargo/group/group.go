package group

import (
	"fmt"
	"github.com/monochromegane/cargo/cargo/option"
	"os/exec"
	"strings"
)

type Group struct {
	From   string
	Option option.Option
}

func (self *Group) GroupBy() map[int][]string {
	if self.Option.Group == "go-package" {
		return self.ByGoPackage()
	}
	return make(map[int][]string)
}

func (self *Group) ByGoPackage() map[int][]string {
	opt := self.Option
	command := exec.Command(
		"docker",
		"run",
		"-v",
		self.From+":"+opt.Dest,
		opt.Image,
		"/go/go/bin/go",
		"list",
		opt.GoPackage+"/...",
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