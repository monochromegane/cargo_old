package command

import (
	"os/exec"
)

type DockerRunCommand struct {
	Image     string
	SrcVolume string
	DstVolume string
	Cmd       []string
}

func (self *DockerRunCommand) Command() *exec.Cmd {

	cmd := []string{"run"}
	if len(self.SrcVolume) > 0 && len(self.DstVolume) > 0 {
		cmd = append(cmd, []string{
			"-v",
			self.SrcVolume + ":" + self.DstVolume}...,
		)
	}
	cmd = append(cmd, self.Image)
	cmd = append(cmd, self.Cmd...)

	return exec.Command("docker", cmd...)
}
