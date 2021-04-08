package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestNew(t *testing.T) {
	var que *Queue
	que = New()
	assert.True(t, que.IsEmpty())
}

func TestQueue_Enqueue(t *testing.T) {
	var que *Queue
	que = New()
	que.Enqueue(1)
	assert.Equal(t, 1, que.Size())
	var i = que.Dequeue().(int)
	assert.Equal(t, 1, i)
	assert.True(t, que.IsEmpty())
}

