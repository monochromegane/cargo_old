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

	group := group.Group{From: asset.CurrentDir(), Option: self.Option}
	groups := group.GroupBy()
	fmt.Printf("%s\n", groups)

	var opt = self.Option
	concurrent := concurrency.Concurrency{}
	concurrent.Run(groups, func(index int, group []string) concurrency.Command {
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
