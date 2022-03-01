package offerII_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_79(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var temp []int
	//visited := make(map[int]bool)
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		//不放
		dfs(index + 1)
		//放入
		temp = append(temp, nums[index])
		dfs(index + 1)
		temp = temp[:len(temp)-1]
	}
	dfs(0)
	return ans
}

//[[1] [1 2] [1 2 3] [1 3] [1 3 3] [2] [2 2] [2 2 3] [2 3] [2 3 3] [3] [3 2] [3 2 3] [3 3] [3 3 3]]
