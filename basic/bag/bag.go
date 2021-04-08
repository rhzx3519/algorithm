package bag

import "github.com/rhzx3519/algorithm/types"

type Bag struct {
	array []types.T
}

func New() *Bag {
	return &Bag{
		array: make([]types.T, 0),
	}
}

func (b *Bag) Add(t types.T) {
	b.array = append(b.array, t)
}

func (b *Bag) Size() int {
	return len(b.array)
}

func (b *Bag) IsEmpty() bool {
	return len(b.array) == 0
}