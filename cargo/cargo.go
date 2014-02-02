package cargo

import (
	"fmt"
        "github.com/monochromegane/cargo/cargo/option"
        "github.com/monochromegane/cargo/cargo/asset"
)

type Cargo struct {
	Option option.Option
}

func (self *Cargo) Run(){
        fmt.Println("Cargo Run!!")
        asset := asset.Asset{self.Option}
        asset.Prepare()
}
