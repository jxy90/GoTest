package easy_test

import "testing"

func Test_majorityElement(t *testing.T) {
	fmt.Println(majorityElement([]int{1, 2, 5, 9, 5, 9, 5, 5, 5}))
}

func majorityElement(nums []int) int {
	cache := make(map[int]int)
	ans := 0
	for _, num := range nums {
		cache[num]++
		if cache[num] > cache[ans] {
			ans = num
		}
	}
	if cache[ans]*2 < len(nums) {
		return -1
	}
	return ans
}
