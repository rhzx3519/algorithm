package merge

import (
	"github.com/rhzx3519/algorithm/iterator"
	"github.com/rhzx3519/algorithm/types"
)

func Sort(a *iterator.Sortable) {
	//n := len(a.List)
	//doSort(a, make([]types.T, n), 0, n-1)
	doSortIteration(a)
}

// 自底向顶
func doSortIteration(a *iterator.Sortable) {
	n := len(a.List)
	aux := make([]types.T, n)
	for sz := 1; sz < n; sz *= 2 {
		for l := 0; l < n - sz; l += sz {
			merge(a, aux, l, l + sz - 1, min(n-1, l+sz+sz-1))
		}
	}
}


// 自顶向底
func doSort(a *iterator.Sortable, aux []types.T, l, r int)  {
	if l >= r {
		return
	}
	var mid = l + (r-l)/2
	doSort(a, aux, l, mid)
	doSort(a, aux, mid+1, r)
	merge(a, aux, l, mid, r)
}

func merge(a *iterator.Sortable, aux []types.T, l, mid, r int) {
	var i, j = l, mid + 1
	for k := l; k <= r; k++ {
		aux[k] = a.List[k]
	}
	for k := l; k <= r; k++ {
		if i > mid {
			a.List[k] = aux[j]
			j++
		} else if j > r {
			a.List[k] = aux[i]
			i++
		} else if a.Cmp(aux[i], aux[j]) < 0 {
			a.List[k] = aux[i]
			i++
		} else {
			a.List[k] = aux[j]
			j++
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
