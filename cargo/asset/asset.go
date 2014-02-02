package asset

import (
	"github.com/monochromegane/cargo/cargo/option"
	"os/exec"
	"path/filepath"
	"time"
)

type Asset struct {
	Option    option.Option
	timestamp string
}

func (self *Asset) Prepare() {
	self.timestamp = string(time.Now().Unix())

}

func (self *Asset) CurrentDir() string {
	return filepath.Join(self.baseDir(), "current")
}

func (self *Asset) WorkDir(index int) string {
	return filepath.Join(self.baseDir(), self.timestamp, string(index))
}

func (self *Asset) baseDir() string {
	opt := self.Option
	return filepath.Join(opt.WorkDir, opt.Image, opt.User)
}
