package sorts

import (
	"github.com/rhzx3519/algorithm/iterator"
	"github.com/rhzx3519/algorithm/sorts/heap"
	"github.com/rhzx3519/algorithm/sorts/insertion"
	"github.com/rhzx3519/algorithm/sorts/merge"
	"github.com/rhzx3519/algorithm/sorts/priority_queue"
	"github.com/rhzx3519/algorithm/sorts/quick"
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

func TestQuick(t *testing.T) {
	a := &iterator.Sortable{
		List: []types.T{3,1,2,6,4,5,5,6},
		Cmp: IntComparator,
	}

	quick.Sort(a)
	assert.Equal(t, []types.T{1,2,3,4,5,5,6,6}, a.List)
}

func TestPriorityQueue(t *testing.T) {
	q := priority_queue.New([]types.T{3,1,2,6,4,5,5,6}, IntComparator)
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 8, q.Size())

	assert.Equal(t, 6, q.Peek())

	expected := []int{6,6,5,5,4,3,2,1}
	actual := []int{}
	for !q.IsEmpty() {
		t := q.Poll().(int)
		actual = append(actual, t)
	}

	assert.Equal(t, expected, actual)
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())
}

func TestHeap(t *testing.T) {
	a := &iterator.Sortable{
		List: []types.T{3,1,2,6,4,5,5,6},
		Cmp: IntComparator,
	}

	heap.Sort(a)
	assert.Equal(t, []types.T{1,2,3,4,5,5,6,6}, a.List)
}