package symbol_table

import "github.com/rhzx3519/algorithm/types"

// 无序符号表
type SymbolTable interface {
	//New() SymbolTable 		// 创建一张符号表
	Put(k types.K, v types.V)	// 将键值对存入表中（若值为空则将键key 从表中删除）
	Get(k types.K) types.V		// 获取键key 对应的值（若键key 不存在则返回null）
	Delete(k types.K)			// 从表中删去键key（及其对应的值）
	Clear()						// 清空符号表
	Keys() []types.K			// 表中的所有键的集合

	Contains(k types.K) bool	// 键key 在表中是否有对应的值
	IsEmpty() bool				// 表是否为空
	Size() int					// 表中的键值对数量
}

// 有序符号表
type OrderedSymbolTable interface {
	Put(k types.K, v types.V)	// 将键值对存入表中（若值为空则将键key 从表中删除）
	Get(k types.K) types.V		// 获取键key 对应的值（若键key 不存在则返回null）
	Delete(k types.K)			// 从表中删去键key（及其对应的值）
	Clear()						// 清空符号表
	Keys() []types.K			// 表中的所有键的集合
	KeysBetween(l, r types.K) []types.K // [l, r] 之间键的集合

	Contains(k types.K) bool	// 键key 在表中是否有对应的值
	IsEmpty() bool				// 表是否为空
	Size() int					// 表中的键值对数量
	SizeBetween(l, r types.K) int	// [l, r] 之间键的数量

	Min() types.K				// 最小的键
	Max() types.K				// 最大的键
	Floor(k types.K) types.K 	// 小于等于key 的最大键
	Ceiling(k types.K) types.K	// 大于等于key 的最小键
 	Rank(k types.K) int 		// 小于key 的键的数量
 	Select(k int) types.K		// 排名为k的键(k starts from 0)
}
