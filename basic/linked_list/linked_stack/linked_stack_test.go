package linked_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedStack_Size(t *testing.T) {
	var st = New()
	assert.True(t, st.IsEmpty())
	assert.Equal(t, 0, st.size)

	st.Push(1)
	assert.False(t, st.IsEmpty())
	assert.Equal(t, 1, st.size)

	st.Push(2)
	assert.Equal(t, 2, st.size)

	assert.Equal(t, 2, st.Pop().(int))
	assert.Equal(t, 1, st.size)

	assert.Equal(t, 1, st.Pop().(int))
	assert.Equal(t, 0, st.size)
	assert.True(t, st.IsEmpty())

	st.Push(3)
	assert.False(t, st.IsEmpty())
	assert.Equal(t, 1, st.size)
}
