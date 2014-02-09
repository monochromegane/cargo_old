package cargo

import (
	"fmt"
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

	fmt.Println("Cargo Run!!")
	asset := asset.Asset{Option: self.Option}
	err := asset.Prepare()
	fmt.Printf("%s\n", err)

        groups := group.GroupBy(asset.CurrentDir(), self.Option).GroupBy()
	fmt.Printf("%s\n", groups)

	var opt = self.Option
	concurrency.Run(groups, func(index int, group []string) concurrency.Commander {
		return docker.RunCommand(
			opt.Image,
			docker.RunOption{
				SrcVolume: asset.WorkDirWithIndex(index),
				DstVolume: opt.Dest,
			},
			append(strings.Split(opt.Command, " "), group...),
		)
	})

}
