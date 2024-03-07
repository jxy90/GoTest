package Dynamic

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

//300. 最长递增子序列
//https://leetcode.cn/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}
	res := 0
	for i := range dp {
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}

//354. 俄罗斯套娃信封问题
//https://leetcode.cn/problems/russian-doll-envelopes/description/
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 1 {
		return 1
	}
	//对数组升序排序，获取升序的子串
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] < envelopes[j][1])
	})
	n := len(envelopes)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}
	res := 0
	for _, i := range dp {
		res = int(math.Max(float64(res), float64(i)))
	}
	return res
}

func Test_1(t *testing.T) {
	//fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))

	fmt.Println(maxEnvelopes([][]int{{1, 1}}))
	fmt.Println(maxEnvelopes([][]int{{4, 5}, {6, 7}, {2, 3}}))
}
