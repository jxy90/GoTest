package middle_test

import (
	"fmt"
	"testing"
)

func Test_executeInstructions(t *testing.T) {
	//fmt.Println(executeInstructions(3, []int{1, 0}, "RRDDLU"))
	fmt.Println(executeInstructions(3, []int{0, 1}, "RRDDLU"))
	fmt.Println(executeInstructions(2, []int{1, 1}, "LURD"))
}

func executeInstructions(n int, startPos []int, s string) []int {
	//记录移动方式
	cache := map[byte]int{'R': 0, 'D': 1, 'L': 2, 'U': 3}
	options := []int{0, 1, 0, -1, 0}

	var dfs func(index, step int, cur []int) int
	dfs = func(index, step int, cur []int) int {
		if index == len(s) {
			return step
		}
		option := cache[s[index]]
		next := []int{cur[0] + options[option], cur[1] + options[option+1]}
		if next[0] >= 0 && next[0] < n && next[1] >= 0 && next[1] < n {
			return dfs(index+1, step+1, next)
		} else {
			return step
		}
	}
	//保存结果
	ans := make([]int, len(s))
	for i := range s {
		temp := dfs(i, 0, startPos)
		ans[i] = temp
	}
	return ans
}
