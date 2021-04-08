package sorts

import (
	"github.com/rhzx3519/algorithm/iterator"
	"github.com/rhzx3519/algorithm/sorts/insertion"
	"github.com/rhzx3519/algorithm/sorts/merge"
	"github.com/rhzx3519/algorithm/sorts/selection"
	"github.com/rhzx3519/algorithm/sorts/shell"
	"github.com/rhzx3519/algorithm/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	IntComparator = func(left, right types.T) int {
		a, b := left.(int), right.(int)
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
)

func TestSelection(t *testing.T) {
	a := &iterator.Sortable{
		List: []types.T{3,1,2},
		Cmp: IntComparator,
	}

	selection.Sort(a)
	assert.Equal(t, []types.T{1,2,3}, a.List)
}

func TestInsertion(t *testing.T) {
	a := &iterator.Sortable{
		List: []types.T{3,1,2},
		Cmp: IntComparator,
	}

	insertion.Sort(a)
	assert.Equal(t, []types.T{1,2,3}, a.List)
}

func TestShell(t *testing.T) {
	a := &iterator.Sortable{
		List: []types.T{3,1,2,6,4,5,5,6},
		Cmp: IntComparator,
	}

	shell.Sort(a)
	assert.Equal(t, []types.T{1,2,3,4,5,5,6,6}, a.List)
}

func TestMerge(t *testing.T) {
	a := &iterator.Sortable{
		List: []types.T{3,1,2,6,4,5,5,6},
		Cmp: IntComparator,
	}

	merge.Sort(a)
	assert.Equal(t, []types.T{1,2,3,4,5,5,6,6}, a.List)
}