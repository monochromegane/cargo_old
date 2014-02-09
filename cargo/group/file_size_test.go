package group

import (
	"github.com/monochromegane/cargo/cargo/group"
	"testing"
)

type Assert struct {
	Number, MinSize int
}

func TestMinimum(t *testing.T) {
	sizes := []int{10, 20, 30, 40, 50}
	asserts := []Assert{
		Assert{1, 150},
		Assert{2, 60},
		Assert{3, 30},
	}

	for _, assert := range asserts {
		groups := group.NewGroups(assert.Number)
		for _, size := range sizes {
			min := groups.Minimum()
			min.TotalSize = min.TotalSize + size
		}
		if groups.Minimum().TotalSize != assert.MinSize {
			t.Errorf("When group number is %d, total size should equal %d.", assert.Number, assert.MinSize)
		}
	}
}
