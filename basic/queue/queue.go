package queue

import (
	"github.com/rhzx3519/algorithm/types"
)

type Queue struct {
	array []types.T
}

func New() *Queue {
	return &Queue{
		array: make([]types.T, 0),
	}
}

func (q *Queue) Enqueue(t types.T) {
	q.array = append(q.array, t)
}

func (q *Queue) Dequeue() (r types.T) {
	if q.IsEmpty() {
		return
	}
	r = q.array[0]
	q.array[0] = nil
	q.array = q.array[1:]
	return
}

func (q *Queue) IsEmpty() bool {
	return len(q.array) == 0
}

func (q *Queue) Size() int {
	return len(q.array)
}