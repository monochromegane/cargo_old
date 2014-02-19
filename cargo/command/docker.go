package command

import (
	"os/exec"
)

type DockerRunCommand struct {
	Image           string
	HostVolume      string
	ContainerVolume string
	Cmd             []string
	WorkDir         string
}

func (self *DockerRunCommand) Command() *exec.Cmd {

	cmd := []string{"run"}
	if len(self.HostVolume) > 0 && len(self.ContainerVolume) > 0 {
		cmd = append(cmd, []string{
			"-v",
			self.HostVolume + ":" + self.ContainerVolume}...,
		)
	}
	if len(self.WorkDir) > 0 {
		cmd = append(cmd, []string{
			"-w",
			self.WorkDir}...,
		)
	}
	cmd = append(cmd, self.Image)
	cmd = append(cmd, self.Cmd...)

	return exec.Command("docker", cmd...)
}
