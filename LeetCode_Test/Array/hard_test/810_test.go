package hard_test

import "testing"

func Test_xorGame(t *testing.T) {
	fmt.Println(xorGame([]int{1, 1, 2}))
}

func xorGame(nums []int) bool {
	sum := 0
	for i := range nums {
		sum ^= nums[i]
	}
	return sum == 0 || len(nums)%2 == 0
}
