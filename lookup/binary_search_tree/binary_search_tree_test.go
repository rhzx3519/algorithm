package binary_search_tree

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

func TestBST_Put(t *testing.T) {
	var bst = New()
	assert.Equal(t, 0, bst.Size())
	assert.True(t, bst.IsEmpty())

	bst.Put(integer(1), "1")
	assert.Equal(t, 1, bst.Size())
	assert.False(t, bst.IsEmpty())

	bst.Put(integer(2), "2")
	assert.Equal(t, 2, bst.Size())
}

func TestBST_Get(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(3), "3")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")
	assert.Equal(t, 5, bst.Size())

	assert.Equal(t, "3", bst.Get(integer(3)).(string))
	assert.Equal(t, "-1", bst.Get(integer(-1)).(string))
	assert.Equal(t, "0", bst.Get(integer(0)).(string))
	assert.Equal(t, "1", bst.Get(integer(1)).(string))
	assert.Equal(t, "2", bst.Get(integer(2)).(string))

	assert.Nil(t, bst.Get(integer(10)))
}

func TestBST_Contains(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(-1), "-1")
	assert.True(t, bst.Contains(integer(1)))
	assert.True(t, bst.Contains(integer(-1)))
	assert.False(t, bst.Contains(integer(13)))
}

func TestBST_Clear(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(-1), "-1")
	assert.Equal(t, 2, bst.Size())
	assert.False(t, bst.IsEmpty())

	bst.Clear()
	assert.Equal(t, 0, bst.Size())
	assert.True(t, bst.IsEmpty())
}

func TestBST_Delete(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(3), "3")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")
	assert.Equal(t, 5, bst.Size())

	bst.Delete(integer(3))
	assert.Equal(t, 4, bst.Size())
	assert.False(t, bst.Contains(integer(3)))

	bst.Delete(integer(-1))
	assert.Equal(t, 3, bst.Size())
	assert.False(t, bst.Contains(integer(-1)))

	bst.Delete(integer(1))
	assert.Equal(t, 2, bst.Size())
	assert.False(t, bst.Contains(integer(1)))

	bst.Delete(integer(10))
	assert.Equal(t, 2, bst.Size())

	bst.Delete(integer(2))
	assert.Equal(t, 1, bst.Size())
	assert.False(t, bst.Contains(integer(2)))

	bst.Delete(integer(0))
	assert.Equal(t, 0, bst.Size())
	assert.False(t, bst.Contains(integer(0)))
	assert.True(t, bst.IsEmpty())
}

func TestBST_Ceiling(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(5), "3")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")

	assert.Equal(t, integer(5), bst.Ceiling(integer(3)))
	assert.Equal(t, integer(2), bst.Ceiling(integer(2)))
	assert.Equal(t, integer(0), bst.Ceiling(integer(0)))
	assert.Equal(t, integer(-1), bst.Ceiling(integer(-2)))

	assert.Nil(t, bst.Ceiling(integer(6)))
}

func TestBST_Floor(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(5), "5")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")

	assert.Equal(t, integer(1), bst.Floor(integer(1)))
	assert.Equal(t, integer(2), bst.Floor(integer(3)))
	assert.Equal(t, integer(5), bst.Floor(integer(5)))
	assert.Equal(t, integer(5), bst.Floor(integer(10)))

	assert.Equal(t, integer(-1), bst.Floor(integer(-1)))
	assert.Nil(t, bst.Floor(integer(-2)))
}

func TestBST_Select(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(5), "5")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")

	assert.Equal(t, integer(-1), bst.Select(0))
	assert.Equal(t, integer(0), bst.Select(1))
	assert.Equal(t, integer(1), bst.Select(2))
	assert.Equal(t, integer(2), bst.Select(3))
	assert.Equal(t, integer(5), bst.Select(4))

	assert.Nil(t, bst.Select(-1))
	assert.Nil(t, bst.Select(5))
}

func TestBST_Rank(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(5), "5")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")

	assert.Equal(t, 0, bst.Rank(integer(-1)))
	assert.Equal(t, 1, bst.Rank(integer(0)))
	assert.Equal(t, 2, bst.Rank(integer(1)))
	assert.Equal(t, 3, bst.Rank(integer(2)))
	assert.Equal(t, 4, bst.Rank(integer(3)))
	assert.Equal(t, 4, bst.Rank(integer(4)))
	assert.Equal(t, 4, bst.Rank(integer(5)))
	assert.Equal(t, 5, bst.Rank(integer(10)))
	assert.Equal(t, 5, bst.Rank(integer(110)))
}

func TestBST_Keys(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(5), "5")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")

	assert.Equal(t, []types.K{integer(-1), integer(0), integer(1), integer(2), integer(5)}, bst.Keys())
}

func TestBST_KeysBetween(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(5), "5")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")

	assert.Equal(t, []types.K{}, bst.KeysBetween(integer(-10), integer(-3)))
	assert.Equal(t, []types.K{integer(-1)}, bst.KeysBetween(integer(-10), integer(-1)))
	assert.Equal(t, []types.K{integer(0), integer(1)}, bst.KeysBetween(integer(0), integer(1)))
	assert.Equal(t, []types.K{integer(0), integer(1), integer(2)}, bst.KeysBetween(integer(0), integer(4)))
	assert.Equal(t, []types.K{integer(0), integer(1), integer(2), integer(5)},
					bst.KeysBetween(integer(0), integer(10)))
}

func TestBST_Min(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(5), "5")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")

	assert.Equal(t, integer(-1), bst.Min())

	bst.Clear()
	assert.Nil(t, bst.Min())
}

func TestBST_Max(t *testing.T) {
	var bst = New()
	bst.Put(integer(1), "1")
	bst.Put(integer(2), "2")
	bst.Put(integer(5), "5")
	bst.Put(integer(-1), "-1")
	bst.Put(integer(0), "0")

	assert.Equal(t, integer(5), bst.Max())

	bst.Clear()
	assert.Nil(t, bst.Max())
}


