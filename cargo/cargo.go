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

func (self *Cargo) Run() bool {

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
		resultBefore, err := before.Command().CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to run before all.\n%s\n%s\n", resultBefore, err)
			return false
		}
	}

	err := asset.Prepare()
	if err != nil {
		fmt.Println("Failed to prepare assets.")
		return false
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
		fmt.Printf("\n========== [concurrency %d] ==========\n%s\n", index, result)
		if err != nil {
			fmt.Printf("Failed to run tests.\n%s\n", err)
		}
		os.RemoveAll(asset.WorkDirWithIndex(index))
		return true
	})
	return true
}

func (self *Cargo) printDebug(log ...interface{}) {
	if self.Option.Debug {
		fmt.Printf("DEBUG: %s\n", log)
	}
}
