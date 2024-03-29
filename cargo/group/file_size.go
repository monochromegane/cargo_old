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
	filter := self.Option.GetFilter()
	filepath.Walk(filepath.Join(self.From, self.Option.Target), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filter != nil && !filter.MatchString(info.Name()) {
			return nil
		}
		min := groups.Minimum()
		min.TotalSize = min.TotalSize + int(info.Size())
		rel, _ := filepath.Rel(self.From, path)
		min.Files = append(min.Files, rel)
		return nil

	})
	return groups.Convert()
}
