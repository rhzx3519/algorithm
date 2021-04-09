package iterator

import "github.com/rhzx3519/algorithm/types"

type iterator interface {
	GetSizeIfKnown() int64
	HasNext() bool
	Next() types.T
}

// region iterable

type iterable interface {
	iter() iterator
}

// end region iterable

// region Sortable
// Sortable use types.Comparator to sort []types.T 可以使用指定的 cmp 比较器对 list 进行排序
// see sort.Interface
type Sortable struct {
	List []types.T
	Cmp types.Comparator
}
// Len is the number of elements in the collection.
func (a *Sortable) Len() int {
	return len(a.List)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (a *Sortable) Less(i, j int) bool {
	return a.Cmp(a.List[i], a.List[j]) < 0
}

// Swap swaps the elements with indexes i and j.
func (a *Sortable) Swap(i, j int) {
	a.List[i], a.List[j] = a.List[j], a.List[i]
}

//end region Sortable
