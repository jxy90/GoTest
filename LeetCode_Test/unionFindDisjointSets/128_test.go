package unionFindDisjointSets

import (
	"github.com/jxy90/GoTest/LeetCode_Test/Struct"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

//循环
func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	memo := make(map[int]bool, 0)
	for i := range nums {
		memo[nums[i]] = true
	}
	ans := 1
	for key, _ := range memo {
		if _, ok := memo[key-1]; ok {
			continue
		}
		count := 0
		for memo[key] {
			key++
			count++
		}
		if count > ans {
			ans = count
		}
	}
	return ans
}

//并查集
func longestConsecutive0(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	memo := make(map[int]int, 0)
	for i := range nums {
		memo[nums[i]] = i
	}
	u := Struct.ConstructorUnionFindHeight(n)
	for value, index := range memo {
		if _, ok := memo[value+1]; !ok {
			continue
		}
		if u.Same(index, memo[value+1]) {
			continue
		}
		u.Union(index, memo[value+1])
	}
	ans := 0
	for _, value := range u.Height {
		if value > ans {
			ans = value
		}
	}
	return ans
}
