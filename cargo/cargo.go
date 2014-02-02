package cargo

import (
	"fmt"
	"github.com/monochromegane/cargo/cargo/asset"
	"github.com/monochromegane/cargo/cargo/option"
)

type Cargo struct {
	Option option.Option
}

func (self *Cargo) Run() {
	fmt.Println("Cargo Run!!")
	asset := asset.Asset{Option: self.Option}
        err := asset.Prepare()
        fmt.Printf("%s\n", err)

}
