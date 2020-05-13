package middle_test

import "testing"

var his map[int]int

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	his = map[int]int{}
	return validateBinaryTreeNodesHelper(n, leftChild, rightChild)
}
func validateBinaryTreeNodesHelper(n int, leftChild []int, rightChild []int) bool {
	_count := n - 1
	for _, value := range leftChild {
		if value != -1 {
			if his[value] == 0 {
				his[0] = 1
				his[value] = 1
			} else {
				return false
			}
			_count--
		}
	}
	for _, value := range rightChild {
		if value != -1 {
			if his[value] == 0 {
				his[0] = 1
				his[value] = 1
			} else {
				return false
			}
			_count--
		}
	}
	return _count == 0
}

func Test_validateBinaryTreeNodes(t *testing.T) {
	left := []int{1, 2, 0, -1}
	right := []int{-1, -1, -1, -1}
	validateBinaryTreeNodes(4, left, right)
}
