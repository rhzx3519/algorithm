package quick

import (
	"github.com/rhzx3519/algorithm/iterator"
)

func Sort(a *iterator.Sortable) {
	n := len(a.List)
	doSort(a, 0, n-1)
}

func doSort(a *iterator.Sortable, l, r int) {
	if l >= r {
		return
	}
	p := partition2(a, l, r)
	doSort(a, l, p)
	doSort(a, p+1, r)
}


// 根据key=a.List[l]的值，将a.List[l:r+1]分成两部分，
// 小于key的元素放在前半部分，大于key的元素放在后半部分

// 顺序遍历，适用于链表排序
func partition1(a *iterator.Sortable, l, r int) int {
	var i, j = l, l
	for i <= r {
		if a.Less(i, l) {
			j++
			a.Swap(i, j)
		}
		i++
	}
	a.Swap(l, j)
	return j
}

func partition2(a *iterator.Sortable, l, r int) int {
	key := a.List[l]
	for l < r {
		for l < r && a.Cmp(key, a.List[r]) <= 0 {
			r--
		}
		a.List[l] = a.List[r]
		for l < r && a.Cmp(key, a.List[l]) >= 0 {
			l++
		}
		a.List[r] = a.List[l]
	}
	a.List[l] = key
	return l
}