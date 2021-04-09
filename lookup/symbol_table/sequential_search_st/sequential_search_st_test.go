package sequential_search_st

import (
	"github.com/rhzx3519/algorithm/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

type integer int

func (i integer) CompareTo(o types.T) int {
	var j = o.(integer)
	if i > j {
		return 1
	} else if i < j {
		return -1
	} else {
		return 0
	}
}

func TestSequentialSearchSt_all(t *testing.T) {
	st := New()
	assert.True(t, st.IsEmpty())
	assert.Equal(t, 0, st.Size())

	st.Put(integer(1), "2")

	assert.False(t, st.IsEmpty())
	assert.Equal(t, 1, st.Size())

	assert.Equal(t, "2", st.Get(integer(1)))

	st.Put(integer(1), "3")
	assert.Equal(t, 1, st.Size())
	assert.Equal(t, "3", st.Get(integer(1)))
}

func TestSequentialSearchSt_Delete(t *testing.T) {
	st := New()

	st.Put(integer(1), "2")
	assert.Equal(t, 1, st.Size())

	st.Delete(integer(1))
	assert.True(t, st.IsEmpty())
	assert.Equal(t, 0, st.Size())

	st.Put(integer(2), "3")
	assert.Equal(t, 1, st.Size())

	st.Put(integer(2), nil)
	assert.True(t, st.IsEmpty())
	assert.Equal(t, 0, st.Size())
}

func TestSequentialSearchSt_Contains(t *testing.T) {
	st := New()

	st.Put(integer(1), "2")
	assert.Equal(t, 1, st.Size())

	assert.True(t, st.Contains(integer(1)))
	assert.False(t, st.Contains(integer(2)))
}

func TestSequentialSearchSt_Keys(t *testing.T) {
	st := New()

	st.Put(integer(1), "1")
	st.Put(integer(2), "2")

	var keys = st.Keys();
	assert.Equal(t, []types.K{integer(1), integer(2)}, keys)
}

func TestSequentialSearchSt_Clear(t *testing.T) {
	st := New()
	st.Put(integer(1), "1")
	st.Put(integer(2), "2")
	assert.Equal(t, 2, st.Size())

	st.Clear()
	assert.Equal(t, 0, st.Size())
	assert.True(t, st.IsEmpty())
}

