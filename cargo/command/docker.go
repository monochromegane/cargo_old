package command

import (
	"os/exec"
)

type DockerRunCommand struct {
	Image           string
	HostVolume      string
	ContainerVolume string
	Cmd             []string
}

func (self *DockerRunCommand) Command() *exec.Cmd {

	cmd := []string{"run"}
	if len(self.HostVolume) > 0 && len(self.ContainerVolume) > 0 {
		cmd = append(cmd, []string{
			"-v",
			self.HostVolume + ":" + self.ContainerVolume}...,
		)
	}
	cmd = append(cmd, self.Image)
	cmd = append(cmd, self.Cmd...)

	return exec.Command("docker", cmd...)
}
