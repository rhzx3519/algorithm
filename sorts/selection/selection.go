package selection

import (
	"github.com/rhzx3519/algorithm/iterator"
)

// ιζ©ζεΊ
// time: O(n^2)
func Sort(a *iterator.Sortable) {
	n := len(a.List)
	var min int
	for i := 0; i < n; i++ {
		min = i
		for j := i+1; j < n; j++ {
			if a.Less(j, min) {
				min = j
			}
		}
		a.Swap(i, min)
	}
}
