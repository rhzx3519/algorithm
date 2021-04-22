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
