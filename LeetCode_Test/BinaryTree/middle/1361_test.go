package middle_test

import "testing"

var his map[int]int

func ValidateBinaryTreeNodes(n int, LeftChild []int, RightChild []int) bool {
	his = map[int]int{}
	return ValidateBinaryTreeNodesHelper(n, LeftChild, RightChild)
}
func ValidateBinaryTreeNodesHelper(n int, LeftChild []int, RightChild []int) bool {
	_count := n - 1
	for _, Value := range LeftChild {
		if Value != -1 {
			if his[Value] == 0 {
				his[0] = 1
				his[Value] = 1
			} else {
				return false
			}
			_count--
		}
	}
	for _, Value := range RightChild {
		if Value != -1 {
			if his[Value] == 0 {
				his[0] = 1
				his[Value] = 1
			} else {
				return false
			}
			_count--
		}
	}
	return _count == 0
}

func Test_ValidateBinaryTreeNodes(t *testing.T) {
	Left := []int{1, 2, 0, -1}
	Right := []int{-1, -1, -1, -1}
	ValidateBinaryTreeNodes(4, Left, Right)
}
