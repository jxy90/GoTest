package hard_test

import (
	"fmt"
	"testing"
)

func Test_longestSubsequence(t *testing.T) {
	fmt.Println(longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}

func longestSubsequence(arr []int, difference int) int {
	ans := 0
	f := map[int]int{}
	//从左向右遍历
	//f 记录当前值在以difference等差队列中的位置
	//结果即为其中出现的最大值
	for _, item := range arr {
		f[item] = f[item-difference] + 1
		if f[item] > ans {
			ans = f[item]
		}
	}
	return ans
}
