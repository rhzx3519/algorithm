package binary_search_tree

import (
	"github.com/rhzx3519/algorithm/lookup/symbol_table"
	"github.com/rhzx3519/algorithm/types"
)

type node struct {
	key 		types.K		// 键
	value 		types.V		// 值
	left, right *node
	N 			int			// 以该结点为根的子树中的结点总数
}

type BST struct {
	root *node
}

func New() symbol_table.OrderedSymbolTable {
	return new(BST)
}

// -------------------------------------------------------------------------------------------
// BST public methods

func (b *BST) Put(k types.K, v types.V) {
	if v == nil {
		b.Delete(k)
		return
	}
	b.root = insert(b.root, k, v)
}

func (b *BST) Get(k types.K) types.V {
	return get(b.root, k)
}

func (b *BST) Delete(k types.K) {
	b.root = delete(b.root, k)
}

func (b *BST) Clear() {
	b.root = nil
}

func (b *BST) Keys() []types.K {
	return b.KeysBetween(b.Min(), b.Max())
}

func (b *BST) KeysBetween(l, r types.K) []types.K {
	var res = make([]types.K, 0)
	keys(b.root, l, r, &res)
	return res
}

func (b *BST) Contains(k types.K) bool {
	return b.Get(k) != nil
}

func (b *BST) IsEmpty() bool {
	return b.Size() == 0
}

func (b *BST) Size() int {
	return size(b.root)
}

func (b *BST) SizeBetween(l, r types.K) int {
	return len(b.KeysBetween(l, r))
}

func (b *BST) Min() types.K {
	node := min(b.root)
	if node != nil {
		return node.key
	}
	return nil
}

func (b *BST) Max() types.K {
	node := max(b.root)
	if node != nil {
		return node.key
	}
	return nil
}

func (b *BST) Floor(k types.K) types.K {
	node := floor(b.root, k)
	if node != nil {
		return node.key
	}
	return nil
}

func (b *BST) Ceiling(k types.K) types.K {
	node := ceiling(b.root, k)
	if node != nil {
		return node.key
	}
	return nil
}

func (b *BST) Rank(k types.K) int {
	return rank(b.root, k)
}

func (b *BST) Select(k int) types.K {
	 node := select_(b.root, k)
	 if node != nil {
	 	return node.key
	 }
	 return nil
}

// -------------------------------------------------------------------------------------------
// private methods

func size(root *node) int {
	if root == nil {
		return 0
	}
	return root.N
}

func get(root *node, k types.K) types.V {
	if root == nil {
		return nil
	}
	if root.key.CompareTo(k) == 0 {
		return root.value
	} else if root.key.CompareTo(k) > 0 {
		return get(root.left, k)
	} else {
		return get(root.right, k)
	}
}

func delete(root *node, k types.K) *node {
	if root == nil {
		return nil
	}

	if root.key.CompareTo(k) < 0 {
		root.right = delete(root.right, k)
	} else if root.key.CompareTo(k) > 0 {
		root.left = delete(root.left, k)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			var tmp = root
			//root = min(root.right)
			//root .right = deleteMin(root.right)
			//root.left = tmp.left
			root = max(root.left)
			root.left = deleteMax(root.left)
			root.right = tmp.right
		}
	}
	root.N = size(root.left) + size(root.right) + 1
	return root
}

func min(root *node) *node {
	if root == nil {
		return root
	}
	if root.left == nil {
		return root
	}
	return min(root.left)
}

func max(root *node) *node {
	if root == nil {
		return root
	}
	if root.right == nil {
		return root
	}
	return max(root.right)
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

func deleteMin(root *node) *node {
	if root == nil {
		return nil
	}
	if root.left == nil {
		return root.right
	}
	root.left = deleteMin(root.left)
	root.N = size(root.left) + size(root.right) + 1
	return root
}

func deleteMax(root *node) *node {
	if root == nil {
		return nil
	}
	if root.right == nil {
		return root.left
	}
	root.right = deleteMax(root.right)
	root.N = size(root.left) + size(root.right) + 1
	return root
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

func insert(root *node, k types.K, v types.V) *node {
	if root == nil {
		root = &node{
			key: k,
			value: v,
			N: 1,
		}
		return root
	}
	var cmp = root.key.CompareTo(k)
	if cmp > 0 {
		root.left = insert(root.left, k, v)
	} else if cmp < 0 {
		root.right = insert(root.right, k, v)
	} else {
		root.value = v
	}
	root.N = size(root.left) + size(root.right) + 1
	return root
}

