package red_black_tree

import (
	"github.com/rhzx3519/algorithm/types"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRedBlackTree_Put(t *testing.T) {
	var rbt = &RedBlackTree{}
	rbt.Put(types.Integer(0), "0")
	rbt.Put(types.Integer(3), "3")
	rbt.Put(types.Integer(7), "7")
	rbt.Put(types.Integer(5), "5")
	rbt.Put(types.Integer(-5), "-5")
	rbt.Put(types.Integer(-15), "-15")
}

func TestRedBlackTree_Delete(t *testing.T) {
	var rbt = getRbt()
	assert.Equal(t, 6, rbt.Size())

	rbt.Delete(types.Integer(0))
	assert.Equal(t, 5, rbt.Size())
	assert.Nil(t, rbt.Get(types.Integer(0)))

	rbt.Delete(types.Integer(math.MaxInt32))
	assert.Equal(t, 5, rbt.Size())

}

func TestRedBlackTree_LowerBound(t *testing.T) {
	var rbt = getRbt()
	assert.Equal(t, types.Integer(0), rbt.LowerBound(types.Integer(0)))
	assert.Equal(t, types.Integer(3), rbt.LowerBound(types.Integer(1)))
	assert.Equal(t, types.Integer(3), rbt.LowerBound(types.Integer(2)))
	assert.Equal(t, types.Integer(3), rbt.LowerBound(types.Integer(3)))
	assert.Equal(t, types.Integer(5), rbt.LowerBound(types.Integer(4)))

	assert.Nil(t, rbt.LowerBound(types.Integer(8)))

	assert.Equal(t, types.Integer(-15), rbt.LowerBound(types.Integer(-16)))
}

func TestRedBlackTree_UpperBound(t *testing.T) {
	var rbt = getRbt()

	assert.Equal(t, types.Integer(3), rbt.UpperBound(types.Integer(0)))
	assert.Equal(t, types.Integer(3), rbt.UpperBound(types.Integer(1)))
	assert.Equal(t, types.Integer(3), rbt.UpperBound(types.Integer(2)))
	assert.Equal(t, types.Integer(5), rbt.UpperBound(types.Integer(3)))
	assert.Equal(t, nil, rbt.UpperBound(types.Integer(7)))

	assert.Equal(t, types.Integer(-15), rbt.UpperBound(types.Integer(-16)))
}


func getRbt() *RedBlackTree {
	var rbt = &RedBlackTree{}
	rbt.Put(types.Integer(0), "0")
	rbt.Put(types.Integer(3), "3")
	rbt.Put(types.Integer(7), "7")
	rbt.Put(types.Integer(5), "5")
	rbt.Put(types.Integer(-5), "-5")
	rbt.Put(types.Integer(-15), "-15")
	return rbt
}
