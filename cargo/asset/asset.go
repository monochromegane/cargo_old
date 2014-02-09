package asset

import (
	"github.com/monochromegane/cargo/cargo/option"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type Asset struct {
	Option    option.Option
	timestamp string
}

func (self *Asset) Prepare() error {
	self.timestamp = string(time.Now().Format("20060102150405"))

	if _, err := os.Stat(self.CurrentDir()); err != nil {
		return err
	}

	// TODO it should execute with go routine.
	os.Mkdir(self.WorkDir(), 0775)
	for i := 0; i < self.Option.Concurrency; i++ {
		command := exec.Command("cp", "-r", self.CurrentDir(), self.WorkDirWithIndex(i))
		err := command.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *Asset) CurrentDir() string {
	return filepath.Join(self.baseDir(), "current")
}

func (self *Asset) WorkDirWithIndex(index int) string {
	return filepath.Join(self.WorkDir(), strconv.Itoa(index))
}

func (self *Asset) WorkDir() string {
	return filepath.Join(self.baseDir(), self.timestamp)
}

func (self *Asset) baseDir() string {
	opt := self.Option
	return filepath.Join(opt.WorkDir, opt.Image, opt.User)
}
