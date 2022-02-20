package easy_test

import (
	"testing"
)

func Test_isOneBitCharacter(t *testing.T) {
	//fmt.Println(kIncreasing([]int{4, 1, 5, 2, 6, 2}, 3))
}

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	index := 0
	for index < n {
		if index == n-1 {
			return true
		}
		if bits[index] == 1 {
			index++
		}
		index++
	}
	return false
}
