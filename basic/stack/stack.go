package stack

import "github.com/rhzx3519/algorithm/types"

type Stack struct {
	array []types.T
}

func New() *Stack {
	return &Stack{
		array: make([]types.T, 0),
	}
}

func (s *Stack) Push(t types.T) {
	s.array = append(s.array, t)
}

func (s *Stack) Pop() types.T {
	if s.IsEmpty() {
		return nil
	}
	r := s.array[len(s.array)-1]
	s.array[len(s.array)-1] = nil
	s.array = s.array[:len(s.array)-1]
	return r
}

func (s *Stack) IsEmpty() bool {
	return len(s.array)	== 0
}

func (s *Stack) Size() int {
	return len(s.array)
}