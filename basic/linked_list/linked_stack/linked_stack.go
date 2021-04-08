package linked_stack

import (
	lk "github.com/rhzx3519/algorithm/basic/linked_list"
	"github.com/rhzx3519/algorithm/types"
)

type LinkedStack struct {
	dummy 	*lk.ListNode
	size 	int
}

func New() *LinkedStack {
	st := &LinkedStack{
		dummy: new(lk.ListNode),
	}
	return st
}

func (st *LinkedStack) Push(t types.T) {
	node := &lk.ListNode{
		Value: t,
	}
	node.Next = st.dummy.Next
	st.dummy.Next = node
	st.size++
}

func (st *LinkedStack) Pop() (r types.T) {
	if st.IsEmpty() {
		return
	}
	r = st.dummy.Next.Value
	st.dummy.Next = st.dummy.Next.Next
	st.size--
	return
}

func (st *LinkedStack) Peek() (t types.T) {
	if st.IsEmpty() {
		return
	}
	return st.dummy.Next.Value
}

func (st *LinkedStack) IsEmpty() bool {
	return st.size == 0
}

func (st *LinkedStack) Size() int {
	return st.size
}


