package cargo

import (
	"fmt"
	"github.com/monochromegane/cargo/cargo/asset"
	"github.com/monochromegane/cargo/cargo/command"
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
	g := group.GroupBy()
	fmt.Printf("%s\n", g)

	var opt = self.Option
	in := make(chan *command.Result)

	for key, value := range g {
		args := []string{
			"run",
			"-v",
			asset.WorkDirWithIndex(key) + ":" + opt.Dest,
			opt.Image,
		}
		args = append(args, strings.Split(opt.Command, " ")...)
		args = append(args, value...)

		command := command.Command{"docker", args}
		go command.Exec(in)
	}

	for i := 0; i < len(g); i++ {
		result := <-in
		fmt.Printf("%s, %s\n", result.Output, result.Err)
	}

}
