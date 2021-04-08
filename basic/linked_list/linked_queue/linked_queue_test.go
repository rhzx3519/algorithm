package linked_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedQueue_Size(t *testing.T) {
	var que *LinkedQueue
	que = New()
	assert.True(t, que.IsEmpty())
	assert.Equal(t, 0, que.size)

	que.Enqueue(1)
	assert.False(t, que.IsEmpty())
	assert.Equal(t, 1, que.size)

	que.Enqueue(2)
	assert.Equal(t, 2, que.size)

	assert.Equal(t, 1, que.Dequeue().(int))
	assert.Equal(t, 1, que.size)
	assert.Equal(t, 2, que.Dequeue().(int))
	assert.Equal(t, 0, que.size)
	assert.True(t, que.IsEmpty())

	que.Enqueue(3)
	assert.Equal(t, 1, que.size)
	assert.False(t, que.IsEmpty())

	assert.Equal(t, 3, que.Dequeue().(int))
	assert.Equal(t, 0, que.size)
	assert.True(t, que.IsEmpty())
}
