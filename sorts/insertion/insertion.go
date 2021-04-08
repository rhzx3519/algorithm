package insertion

import "github.com/rhzx3519/algorithm/iterator"

func Sort(a *iterator.Sortable) {
	n := len(a.List)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if a.Less(j, j-1) {
				a.Swap(j, j-1)
			}
		}
	}
}
