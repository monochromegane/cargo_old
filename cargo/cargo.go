package cargo

import (
	"github.com/monochromegane/cargo/cargo/asset"
	"github.com/monochromegane/cargo/cargo/concurrency"
	"github.com/monochromegane/cargo/cargo/docker"
	"github.com/monochromegane/cargo/cargo/group"
	"github.com/monochromegane/cargo/cargo/option"
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

	concurrency.Run(groups, func(index int, group []string) concurrency.Commander {
		return docker.RunCommand(
			self.Option.Image,
			docker.RunOption{
				SrcVolume: asset.WorkDirWithIndex(index),
				DstVolume: self.Option.Dest,
			},
			append(strings.Split(self.Option.Command, " "), group...),
		)
	})

}
