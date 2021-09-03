package middle_test

import (
	"sort"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	fmt.Println(largestDivisibleSubset([]int{2, 4}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))
}

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	//dp[i]保存前i个位置能被nums[i]整除
	dp := make([]int, n)
	//缓存符合条件的前一个index
	memo := make([]int, n)
	//dp初始化，包含自己
	for i := range dp {
		dp[i] = 1
	}
	maxIndex, maxVal := 0, 1
	for i := 1; i < n; i++ {
		for j := range nums[:i] {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				//j被i整除 && 数量大于上一次的指
				dp[i] = dp[j] + 1
				//保存上一个的index
				memo[i] = j
				if dp[i] > maxVal {
					//记录符合条件的最长数组和记录index
					maxIndex = i
					maxVal = dp[i]
				}
			}
		}
	}
	ans := []int{}
	for maxVal > 0 {
		//倒叙
		ans = append([]int{nums[maxIndex]}, ans...)
		maxIndex = memo[maxIndex]
		maxVal--
	}
	return ans
}
