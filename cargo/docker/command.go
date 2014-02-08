package docker

import (
	"os/exec"
        "fmt"
)

type RunOption struct {
	SrcVolume string
	DstVolume string
}

func RunCommand(image string, opt RunOption, commmands []string) *exec.Cmd {

	args := []string{"run"}
	if len(opt.SrcVolume) > 0 && len(opt.DstVolume) > 0 {
		args = append(args, []string{
			"-v",
			opt.SrcVolume + ":" + opt.DstVolume}...,
		)
	}
	args = append(args, image)
	args = append(args, commmands...)
        fmt.Printf("%s\n", args)

	return exec.Command("docker", args...)
}
