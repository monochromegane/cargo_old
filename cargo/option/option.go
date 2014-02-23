package option

import (
	"regexp"
)

type Option struct {
	Debug       bool   `long:"debug"`
	Image       string `short:"i" long:"image" required:"true" description:"docker image name"`
	Concurrency int    `short:"n" default:"1" default-mask:"-" description:"Number of docker run concurrency"`
	User        string `short:"u" default:"cargo"`
	GroupBy     string `short:"g" default:"file-size"`
	Mount       string `short:"m"`
	Command     string `short:"c"`
	Target      string `short:"t"`
	Filter      string `short:"f" default:""`
	WorkDir     string `short:"w" default:"/tmp/cargo"`
}

func (self *Option) GetFilter() *regexp.Regexp {
	if len(self.Filter) == 0 {
		return nil
	}
	reg, err := regexp.Compile(self.Filter)
	if err != nil {
		return nil
	}
	return reg
}
