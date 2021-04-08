package linked_queue

import (
	lk "github.com/rhzx3519/algorithm/basic/linked_list"
	"github.com/rhzx3519/algorithm/types"
)

type LinkedQueue struct {
	dummy 	*lk.ListNode
	tail 	*lk.ListNode
	size 	int
}

func New() *LinkedQueue {
	var que = &LinkedQueue{
		dummy: new(lk.ListNode),
	}
	que.tail = que.dummy

	return que
}

func (q *LinkedQueue) Enqueue(t types.T)  {
	q.tail.Next = &lk.ListNode{Value: t}
	q.tail = q.tail.Next
	q.size++
}

func (q *LinkedQueue) Dequeue() (r types.T) {
	if q.IsEmpty() {
		return
	}
	r = q.dummy.Next.Value
	q.dummy.Next = q.dummy.Next.Next
	q.size--
	if q.size == 0 {
		q.tail = q.dummy
	}
	return
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedQueue) Size() int {
	return q.size
}
