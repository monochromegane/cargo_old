package cargo

import (
	"fmt"
	"github.com/monochromegane/cargo/cargo/asset"
	"github.com/monochromegane/cargo/cargo/command"
	"github.com/monochromegane/cargo/cargo/concurrency"
	"github.com/monochromegane/cargo/cargo/group"
	"github.com/monochromegane/cargo/cargo/option"
	"os"
	"os/exec"
	"strings"
)

type Cargo struct {
	Option option.Option
}

func (self *Cargo) Run() {

	asset := asset.Asset{Option: self.Option}
	err := asset.Prepare()
	if err != nil {
		return
	}

	groups := group.NewGrouper(asset.CurrentDir(), self.Option).GroupBy()

	concurrency.Run(groups, func(index int, group []string) *exec.Cmd {
		command := command.DockerRunCommand{
			Image:           self.Option.Image,
			HostVolume:      asset.WorkDirWithIndex(index),
			ContainerVolume: self.Option.Mount,
			Cmd:             append(strings.Split(self.Option.Command, " "), group...),
		}
		return command.Command()
	}, func(index int, group []string, result []byte, err error) bool {
		if err != nil {
			return false
		}
		os.RemoveAll(asset.WorkDirWithIndex(index))
		fmt.Printf("%s", result)
		return true
	})

}
