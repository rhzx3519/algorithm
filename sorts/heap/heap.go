package heap

import "github.com/rhzx3519/algorithm/iterator"

// 堆排序, a.List中的元素范围为1~N
func Sort(a *iterator.Sortable) {
	N := len(a.List)
	last := a.List[N-1]
	copy(a.List[1:], a.List)
	a.List = append(a.List, last)

	// 如果一个结点的两个子结点都已经是堆了，那么在该结点上调用sink() 可以将它们变成一个堆。这个过
	// 程会递归地建立起堆的秩序
	// 1 ~ N/2 的子节点堆化
	for k := N/2; k >= 1; k-- {
		sink(a, k, N)
	}

	for N > 1 {
		a.Swap(1, N)
		N--
		sink(a, 1, N)
	}

	a.List = a.List[1:]
}

// 下沉
func sink(a *iterator.Sortable, k, N int) {
	for 2*k <= N {
		j := 2*k
		if j+1 <= N && a.Less(j, j+1) {
			j++
		}
		if a.Less(j, k) {
			break
		}
		a.Swap(j, k)
		k = j
	}
}
