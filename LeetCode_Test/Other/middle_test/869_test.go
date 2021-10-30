package middle_test

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func Test_reorderedPowerOf2(t *testing.T) {
	//fmt.Println(reorderedPowerOf2(46))
	fmt.Println(reorderedPowerOf2(10))
}

func reorderedPowerOf2(n int) (ans bool) {
	nums := []byte(strconv.Itoa(n))
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	//dfs循环所有组合方式
	visited := map[int]bool{}
	var dfs func(index, num int) bool
	dfs = func(index, num int) bool {
		if index == len(nums) {
			fmt.Println(num)
			return num&(num-1) == 0
		}
		for i, char := range nums {
			if (index == 0 && char == '0') || visited[i] || (i > 0 && !visited[i-1] && char == nums[i-1]) {
				continue
			}
			visited[i] = true
			if dfs(index+1, num*10+int(char-'0')) {
				return true
			}
			visited[i] = false
		}
		return false
	}
	return dfs(0, 0)
}
