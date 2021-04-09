package types

type (
	// T indicates any type
	T interface {}
	// R indicates any type
	R interface {}
	// U indicates any type
	U interface {}
	// K indicates a type can be compared
	K Comparable
	// V indicates any type
	V interface {}

	// Comparator is a BiFunction, which two input arguments are the type, and returns a int.
	// if t1 is greater then t2, it returns a positive number;
	// if t1 is less then t2, it returns a negative number; if the two input are equal, it returns 0
	Comparator func(left, right T) int

	Comparable interface {
		CompareTo(o T) int
	}
)

type (
	ListNode struct {
		Value T
		Next *ListNode
	}

	TreeNode struct {
		Value T
		Left *TreeNode
		Right *TreeNode
	}
)