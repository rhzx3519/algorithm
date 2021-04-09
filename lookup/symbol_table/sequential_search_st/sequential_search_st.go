package sequential_search_st

import (
	"github.com/rhzx3519/algorithm/lookup/symbol_table"
	"github.com/rhzx3519/algorithm/types"
)

/*
 使用无序链表实现的符号表
*/

type node struct {
	key types.K
	value types.V
	next *node
}

type SequentialSearchSt struct {
	dummy 	*node
	size	int
}

func New() symbol_table.SymbolTable {
	return &SequentialSearchSt{
		dummy: new(node),
	}
}

func (s *SequentialSearchSt) Put(k types.K, v types.V) {
	if v == nil {
		s.Delete(k)
		return
	}

	var p = s.dummy
	for ; p.next != nil; p = p.next {
		if p.next.key.CompareTo(k) == 0 {
			p.next.value = v
			return
		}
	}
	p.next = &node{
		key: k,
		value: v,
	}
	s.size++
}

func (s *SequentialSearchSt) Get(k types.K) types.V {
	for p := s.dummy; p.next != nil; p = p.next {
		if p.next.key.CompareTo(k) == 0 {
			return p.next.value
		}
	}
	return nil
}

func (s *SequentialSearchSt) Delete(k types.K) {
	for p := s.dummy; p.next != nil; p = p.next {
		if p.next.key.CompareTo(k) == 0 {
			p.next = p.next.next
			s.size--
			return
		}
	}
}

func (s *SequentialSearchSt) Clear() {
	s.dummy.next = nil
	s.size = 0
	return
}

func (s *SequentialSearchSt) Keys() []types.K {
	keys := make([]types.K, 0)
	for p := s.dummy; p.next != nil; p = p.next {
		keys = append(keys, p.next.key)
	}
	return keys
}

func (s *SequentialSearchSt) Contains(k types.K) bool {
	return s.Get(k) != nil
}

func (s *SequentialSearchSt) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SequentialSearchSt) Size() int {
	return s.size
}