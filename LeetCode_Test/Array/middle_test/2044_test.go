package middle_test

import (
	"testing"
)

func Test_2044(t *testing.T) {
	t.Log(countMaxOrSubsets([]int{3, 1}))
	t.Log(countMaxOrSubsets([]int{2, 2, 2}))
	t.Log(countMaxOrSubsets([]int{3, 2, 1, 5}))
}

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	cache := map[int]int{}
	maxVal := 0
	temp := make([]int, 0, n)
	var dfs func(index int)
	//ans := make([][]int, 0)
	dfs = func(index int) {
		if len(temp) > 0 {
			or := temp[0]
			for i := 1; i < len(temp); i++ {
				or = or | temp[i]
			}
			if or > maxVal {
				maxVal = or
			}
			cache[or]++
			//ans = append(ans, append([]int{}, temp...))
		}
		for i := index; i < n; i++ {
			temp = append(temp, nums[i])
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	//fmt.Println(ans)
	return cache[maxVal]
}
