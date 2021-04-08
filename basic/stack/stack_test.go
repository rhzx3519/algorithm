package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Size(t *testing.T) {
	st := New()
	assert.True(t, st.IsEmpty())
	st.Push(1)
	st.Push("1")
	assert.Equal(t, 2, st.Size())
	assert.False(t, st.IsEmpty())

	assert.Equal(t, "1", st.Pop().(string))
	assert.Equal(t, 1, st.Pop().(int))
	assert.True(t, true, st.IsEmpty())
}
