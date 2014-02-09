package group

import (
	"github.com/monochromegane/cargo/cargo/option"
	"os"
	"path/filepath"
)

type FileSize struct {
	From   string
	Option option.Option
}

func NewFileSize(from string, option option.Option) *FileSize {
	return &FileSize{
		From:   from,
		Option: option,
	}
}

type Group struct {
	TotalSize int
	Files     []string
}

type Groups struct {
	Max   int
	Group []*Group
}

func NewGroups(max int) *Groups {
	return &Groups{Max: max}
}

func (self *Groups) Minimum() *Group {
	if self.Max > len(self.Group) {
		group := &Group{}
		self.Group = append(self.Group, group)
		return group
	} else {
		min := &Group{TotalSize: int(^uint(0) >> 1)}
		for _, g := range self.Group {
			if g.TotalSize < min.TotalSize {
				min = g
			}
		}
		return min
	}
}

func (self *Groups) Convert() map[int][]string {
	result := make(map[int][]string)
	for i, g := range self.Group {
		result[i] = g.Files
	}
	return result
}

func (self *FileSize) GroupBy() map[int][]string {

	groups := NewGroups(self.Option.Concurrency)
	filepath.Walk(self.From, func(path string, info os.FileInfo, err error) error {
		min := groups.Minimum()
		min.TotalSize = min.TotalSize + int(info.Size())
		min.Files = append(min.Files, path)
		return nil

	})
	return groups.Convert()
}
