package group

import (
	"github.com/monochromegane/cargo/cargo/option"
)

type Grouper interface {
	GroupBy() map[int][]string
}

func NewGrouper(from string, option option.Option) Grouper {
	switch option.Group {
	case "file-size":
		return NewFileSize(from, option)
	case "go-package":
		return NewGoPackage(from, option)
	default:
		return nil
	}
}
