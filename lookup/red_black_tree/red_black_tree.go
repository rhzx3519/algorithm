/**
红黑树是一种含有红黑结点并能自平衡的二叉查找树。它必须满足下面性质：
性质1：每个节点要么是黑色，要么是红色。
性质2：根节点是黑色。
性质3：每个叶子节点（NIL）是黑色。
性质4：每个红色结点的两个子结点一定都是黑色。
性质5：任意一结点到每个叶子结点的路径都包含数量相同的黑结点。

作者：安卓大叔
链接：https://www.jianshu.com/p/e136ec79235c
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

Insert时维护红黑树的性质:
1. 如果右子结点是红色的而左子结点是黑色的，进行左旋转；
2. 如果左子结点是红色的且它的左子结点也是红色的，进行右旋转；
3. 如果左右子结点均为红色，进行颜色转换
 */
package red_black_tree

import (
	"errors"
	"github.com/rhzx3519/algorithm/lookup/symbol_table"
	"github.com/rhzx3519/algorithm/types"
)

type Color int

const (
	Red Color =	1 << iota
	Black
)

var (
	ErrKeyEmpty = errors.New("Cannot set nil as tree's key.")
	ErrTreeEmpty = errors.New("Cannot operate on an empty tree.")
)

type node struct {
	key    types.K
	value  types.V
	parent *node		// 指向父节点
	left   *node
	right  *node
	size   int
	color  Color		// 从父节点指向改节点的链接是否是红链接
}

type RedBlackTree struct {
	root *node
}

func New() symbol_table.OrderedSymbolTable {
	return new(RedBlackTree)
}

// --------------------------------------------------------------------------------
// RedBlackTree public methods

// 查找key，找到则更新其值，否则为它新建一个结点
func (rbt *RedBlackTree) Put(k types.K, v types.V) {
	if k == nil {
		panic(ErrKeyEmpty)
	}
	if v == nil {
		rbt.Delete(k)
		return
	}
	rbt.root = put(rbt.root, k, v)
	rbt.root.color = Black

	check(rbt)
}

func (rbt *RedBlackTree) Get(k types.K) types.V {
	h := get(rbt.root, k)
	if h == nil {
		return nil
	}
	return h.value
}

func (rbt *RedBlackTree) Delete(k types.K) {
	if k == nil {
		panic(ErrKeyEmpty)
	}
	if !rbt.Contains(k) {
		return
	}

	// if both children of root are black, set root to red
	if !isRed(rbt.root.left) && !isRed(rbt.root.right) {
		rbt.root.color = Red
	}
	rbt.root = delete(rbt.root, k)
	if !rbt.IsEmpty() {
		rbt.root.color = Black
	}
	check(rbt)
}

func (rbt *RedBlackTree) Clear() {
	rbt.root = nil
}

func (rbt *RedBlackTree) Keys() []types.K {
	return rbt.KeysBetween(rbt.Min(), rbt.Max())
}

func (rbt *RedBlackTree) KeysBetween(l, r types.K) []types.K {
	res := []types.K{}
	keys(rbt.root, l, r, &res)
	return res
}

func (rbt *RedBlackTree) Contains(k types.K) bool {
	return rbt.Get(k) != nil
}

func (rbt *RedBlackTree) IsEmpty() bool {
	return rbt.Size() == 0
}

func (rbt *RedBlackTree) Size() int {
	return size(rbt.root)
}

func (rbt *RedBlackTree) SizeBetween(l, r types.K) int {
	return len(rbt.KeysBetween(l, r))
}

func (rbt *RedBlackTree) Min() types.K {
	if rbt.IsEmpty() {
		panic(ErrTreeEmpty)
	}
	return min(rbt.root).key
}

func (rbt *RedBlackTree) Max() types.K {
	if rbt.IsEmpty() {
		panic(ErrTreeEmpty)
	}
	return max(rbt.root).key
}

func (rbt *RedBlackTree) Floor(k types.K) types.K {
	node := floor(rbt.root, k)
	if node == nil {
		return nil
	}
	return node.key
}

func (rbt *RedBlackTree) Ceiling(k types.K) types.K {
	node := ceiling(rbt.root, k)
	if node == nil {
		return nil
	}
	return node.key
}

func (rbt *RedBlackTree) Rank(k types.K) int {
	return rank(rbt.root, k)
}

func (rbt *RedBlackTree) Select(k int) types.K {
	node := select_(rbt.root, k)
	if node != nil {
		return node.key
	}
	return nil
}

// --------------------------------------------------------------------------------
// private methods

func isRed(root *node) bool {
	if root == nil {
		return false
	}
	switch root.color {
	case Red:
		return true
	case Black:
		return false
	default:
		return false
	}
}

// 返回祖父节点
func grandparent(root *node) *node {
	return root.parent.parent
}

// 返回父节点的兄弟节点
func uncle(root *node) *node {
	grand := grandparent(root)
	if grand == nil {
		return nil
	}
	if root.parent == grand.left {
		return grand.right
	} else {
		return grand.right
	}
}

func size(root *node) int {
	if root == nil {
		return 0
	}
	return root.size
}

func get(h *node, key types.K) *node {
	if h == nil {
		return nil
	}
	var cmp = h.key.CompareTo(key)
	if cmp > 0 {
		return get(h.left, key)
	} else if cmp < 0 {
		return get(h.right, key)
	} else {
		return h
	}
}

func put(h *node, key types.K, value types.V) *node {
	if h == nil {
		return &node{
			key:   key,
			value: value,
			size:  1,
			color: Red,
		}
	}

	var cmp = h.key.CompareTo(key)
	if cmp > 0 {
		h.left = put(h.left, key, value)
	} else if cmp < 0 {
		h.right = put(h.right, key, value)
	} else {
		h.value = value
	}

	// 调整红黑树
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}
	h.size = size(h.left) + size(h.right) + 1

	return h
}

/***************************************************************************
 *  Red-black tree deletion.
 ***************************************************************************/

// delete the key-value pair with the minimum key rooted at h
func deleteMin(h *node) *node {
	if h.left == nil {
		return nil
	}

	if !isRed(h.left) && !isRed(h.left.left) {
		h = moveRedLeft(h)
	}
	h.left = deleteMin(h.left)
	return balance(h)
}

// delete the key-value pair with the maximum key rooted at h
func deleteMax(h *node) *node {
	if isRed(h.left) {
		h = rotateRight(h)
	}
	if h.right == nil {
		return nil
	}

	if !isRed(h.right) && !isRed(h.right.left) {
		h = moveRedRight(h)
	}
	h.right = deleteMax(h.right)
	return balance(h)
}

// delete the key-value pair with the given key rooted at h
func delete(h *node, key types.K) *node {
	if key.CompareTo(h.key) < 0 {
		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left = delete(h.left, key)
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		if key.CompareTo(h.key) == 0 && h.right == nil {
			return nil
		}
		if !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		if key.CompareTo(h.key) == 0 {
			var x = min(h.right)
			h.key = x.key
			h.value = x.value
			h.right = deleteMin(h.right)
		} else {
			h.right = delete(h.right, key)
		}
	}

	return balance(h)
}

/***************************************************************************
 *  Red-black tree helper functions.
 ***************************************************************************/

// 左旋，h是子树根节点，x是h的右子节点
func rotateLeft(h *node) *node {
	var x = h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = Red
	x.size = h.size
	h.size = size(h.left) + size(h.right) + 1
	return x
}

// 右旋，h是子树根节点，x是h的左子节点
func rotateRight(h *node) *node {
	var x = h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = Red
	x.size = h.size
	h.size = size(h.left) + size(h.right) + 1
	return x
}

// flip the color of a node and its two children
func flipColors(h *node) {
	if h.color == Red {
		h.color = Black
	} else {
		h.color = Red
	}

	if h.left.color == Red {
		h.left.color = Black
	} else {
		h.left.color = Red
	}

	if h.right.color == Red {
		h.right.color = Black
	} else {
		h.right.color = Red
	}
}

// Assuming that h is red and both h.left and h.left.left
// are black, make h.left or one of its children red.
func moveRedLeft(h *node) *node {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}

// Assuming that h is red and both h.right and h.right.left
// are black, make h.right or one of its children red.
func moveRedRight(h *node) *node {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

// restore red-black tree invariant
func balance(h *node) *node {
	if isRed(h.right) { 	// 右链接是红色
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) { // 连续的红色左链接
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {	// 左右子节点都是红色
		flipColors(h)
	}

	h.size = size(h.left) + size(h.right) + 1
	return h
}

/***************************************************************************
 *  Ordered symbol table methods.
 ***************************************************************************/

func min(x *node) *node {
	if x.left == nil {
		return x
	}
	return min(x.left)
}

func max(x *node) *node {
	if x.right == nil {
		return x
	}
	return max(x.right)
}

func floor(root *node, k types.K) *node {
	if root == nil {
		return nil
	}
	var cmp = root.key.CompareTo(k)
	if cmp == 0 {
		return root
	} else if cmp > 0 {
		return floor(root.left, k)
	} else {
		var tmp = floor(root.right, k)
		if tmp != nil {
			return tmp
		} else {
			return root
		}
	}
}

func ceiling(root *node, k types.K) *node {
	if root == nil {
		return nil
	}
	var cmp = root.key.CompareTo(k)
	if cmp == 0 {
		return root
	} else if cmp < 0 {
		return ceiling(root.right, k)
	} else {
		var tmp = ceiling(root.left, k)
		if tmp != nil {
			return tmp
		} else {
			return root
		}
	}
}

// 返回以root为根结点的子树中小于x.key的键的数量
func rank(root *node, k types.K) int {
	if root == nil {
		return 0
	}
	var cmp = root.key.CompareTo(k)
	if cmp > 0 {
		return rank(root.left, k)
	} else if cmp < 0 {
		return size(root.left) + 1 + rank(root.right, k)
	} else {
		return size(root.left)
	}
}

// 返回排名为k的结点
func select_(root *node, k int) *node {
	if root == nil {
		return root
	}
	var sz = size(root.left)
	if sz > k {
		return select_(root.left, k)
	} else if sz < k {
		return select_(root.right, k - sz - 1)
	} else {
		return root
	}
}


// In order traverse
func keys(root *node, l, r types.K, res *[]types.K) {
	if root == nil {
		return
	}

	var cmpLeft, cmpRight = root.key.CompareTo(l), root.key.CompareTo(r)
	if cmpLeft > 0 {
		keys(root.left, l, r, res)
	}

	if cmpLeft >= 0 && cmpRight <= 0 {
		*res = append(*res, root.key)
	}

	if cmpRight < 0 {
		keys(root.right, l, r, res)
	}
}


/***************************************************************************
 *  Check integrity of red-black tree data structure.
 ***************************************************************************/

func check(rbt *RedBlackTree) {
	if !isBST(rbt) {
		panic("Check BST failed.")
	}
	if !isSizeConsistent(rbt) {
		panic("Check size consistent failed.")
	}
	if !is23(rbt) {
		panic("Check 23 failed.")
	}
	if !isBalanced(rbt) {
		panic("Check balanced failed.")
	}

	if !isRankConsistent(rbt) {
		panic("Check tree node rank failed")
	}
}

// 校验二叉查找树的性质
func isBST(rbt *RedBlackTree) bool {
	var is func(x *node, min, max types.K) bool

	is = func(x *node, min, max types.K) bool {
		if x == nil {
			return true
		}
		if min != nil && x.key.CompareTo(min) <= 0 {
			return false
		}
		if max != nil && x.key.CompareTo(max) >= 0 {
			return false
		}
		return is(x.left, min, x.key) && is(x.right, x.key, max)
	}
	return is(rbt.root, nil, nil)
}

// 校验每个节点的size属性是否正确
func isSizeConsistent(rbt *RedBlackTree) bool {
	var is func(x *node) bool

	is = func(x *node) bool {
		if x == nil {
			return true
		}
		if x.size != size(x.left) + size(x.right) + 1 {
			return false
		}
		return is(x.left) && is(x.right)
	}
	return is(rbt.root)
}

// check that ranks are consistent
func isRankConsistent(rbt *RedBlackTree) bool {
	 for i := 0; i < size(rbt.root); i++ {
		if i != rbt.Rank(rbt.Select(i)) {
			return false
		}
	 }
	 for _, key := range rbt.Keys() {
	 	if key.CompareTo(rbt.Select(rbt.Rank(key))) != 0 {
	 		return false
		}
	 }
	 return true
}

// 校验红黑树没有红色的右子节点，且没有连续红节点(自己是红色、子节点也有红色)
func is23(rbt *RedBlackTree) bool {
	var is func(x *node) bool

	is = func(x *node) bool {
		if x == nil {
			return true
		}
		if isRed(x.right) {
			return false
		}
		if x != rbt.root && isRed(x) && isRed(x.left) {
			return false
		}
		return is(x.left) && is(x.right)
	}
	return is(rbt.root)
}

// do all paths from root to leaf have same number of black edges?
func isBalanced(rbt *RedBlackTree) bool {
	var is func(x *node, black int) bool

	is = func(x *node, black int) bool {
		if x == nil {
			return black == 0
		}
		if !isRed(x) {
			black--
		}
		return is(x.left, black) && is(x.right, black)
	}

	var black int
	var x = rbt.root
	for x != nil {
		if !isRed(x) {
			black++
		}
		x = x.left
	}
	return is(rbt.root, black)
}













