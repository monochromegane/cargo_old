package group

import (
	"github.com/monochromegane/cargo/cargo/option"
)

type Grouper interface {
	GroupBy() map[int][]string
}

func NewGrouper(from string, option option.Option) Grouper {
	if option.Group == "go-package" {
                return NewGoPackage(from, option)
	}
	return nil
}

