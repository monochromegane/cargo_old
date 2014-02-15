package command

import (
	"os/exec"
)

type GoListCommand struct {
	Package string
}

func (self *GoListCommand) Command() *exec.Cmd {
	return exec.Command("go", []string{"list", self.Package + "/..."}...)
}
