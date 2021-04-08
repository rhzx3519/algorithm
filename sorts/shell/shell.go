package shell

import "github.com/rhzx3519/algorithm/iterator"

func Sort(a *iterator.Sortable) {
	n := len(a.List)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			// 将a[i]插入到a[i-h], a[i-2*h], a[i-3*h]... 之中
			for j := i; j >= h; j -= h {
				if a.Less(j, j-h) {
					a.Swap(j, j-h)
				}
			}
		}
		h /= 3
	}
}
