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

	if len(self.Option.BeforeAll) > 0 {
		before := command.DockerRunCommand{
			Image:           self.Option.Image,
			HostVolume:      asset.CurrentDir(),
			ContainerVolume: self.Option.Mount,
			WorkDir:         self.Option.Mount,
			Cmd:             strings.Split(self.Option.BeforeAll, " "),
		}
		self.printDebug(before.Command().Args)
		before.Command().Run()
	}

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
			WorkDir:         self.Option.Mount,
			Cmd:             append(strings.Split(self.Option.Command, " "), group...),
		}
		self.printDebug(command.Command().Args)
		return command.Command()
	}, func(index int, group []string, result []byte, err error) bool {
		self.printDebug(err)
		if err != nil {
			return false
		}
		self.printDebug(result)
		os.RemoveAll(asset.WorkDirWithIndex(index))
		return true
	})

}

func (self *Cargo) printDebug(log ...interface{}) {
	if self.Option.Debug {
		fmt.Printf("DEBUG: %s\n", log)
	}
}
