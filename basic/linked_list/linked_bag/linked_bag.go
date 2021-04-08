package linked_bag

import (
	lk "github.com/rhzx3519/algorithm/basic/linked_list"
	"github.com/rhzx3519/algorithm/types"
)

type LinkedBag struct {
	dummy *lk.ListNode
	size  int
}

func New() *LinkedBag {
	return &LinkedBag{
		dummy: new(lk.ListNode),
	}
}

func (b *LinkedBag) Add(t types.T) {
	node := &lk.ListNode{
		Value: t,
	}
	node.Next = b.dummy.Next
	b.dummy.Next = node
	b.size++
}

func (b *LinkedBag) Size() int {
	return b.size
}

func (b *LinkedBag) IsEmpty() bool {
	return b.size == 0
}
