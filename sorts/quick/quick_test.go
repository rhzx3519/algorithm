package quick

import (
	"github.com/rhzx3519/algorithm/iterator"
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


func TestPartition1(t *testing.T) {
	a := &iterator.Sortable{
		List: []types.T{5,1,5,7},
		Cmp: IntComparator,
	}

	partition1(a, 0, a.Len()-1)
	assert.Equal(t, []types.T{1,5,5,7}, a.List)
}

func TestPartition2(t *testing.T) {
	a := &iterator.Sortable{
		List: []types.T{5,1,5,7},
		Cmp: IntComparator,
	}

	partition2(a, 0, a.Len()-1)
	assert.Equal(t, []types.T{1,5,5,7}, a.List)
}
