package Chapter0_test

import (
	"fmt"
	"sort"
	"testing"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	n := len(nums)

	var helper func(start int)
	helper = func(start int) {
		res = append(res, append([]int{}, track...))
		for i := start; i < n; i++ {
			track = append(track, nums[i])
			helper(i + 1)
			track = track[:len(track)-1]
		}
	}
	helper(0)
	return res
}

//46. 全排列
//https://leetcode.cn/problems/permutations/description/
func permute0(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	visited := make(map[int]struct{})
	n := len(nums)

	var helper func()
	helper = func() {
		if len(track) == n {
			res = append(res, append([]int{}, track...))
		}
		for i := 0; i < n; i++ {
			if _, ok := visited[i]; ok {
				continue
			}
			track = append(track, nums[i])
			visited[i] = struct{}{}
			helper()
			track = track[:len(track)-1]
			delete(visited, i)
		}
	}
	helper()

	return res
}

//90. 子集 II
//https: //leetcode.cn/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	track := make([]int, 0)
	var helper func(start int)
	helper = func(start int) {
		res = append(res, append([]int{}, track...))
		for i := start; i < n; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			if i < n {
				helper(i + 1)
			}
			track = track[:len(track)-1]
		}
	}
	helper(0)

	return res
}

//40. 组合总和 II
//https://leetcode.cn/problems/combination-sum-ii/description/
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	sum := 0
	sort.Ints(candidates)

	var helper func(start int)
	helper = func(start int) {
		if sum == target {
			res = append(res, append([]int{}, track...))
		}
		for i := start; i < len(candidates); i++ {
			if sum > target {
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			track = append(track, candidates[i])
			sum += candidates[i]
			helper(i + 1)
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}
	helper(0)
	return res
}

//47. 全排列 II
//https://leetcode.cn/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	res := make([][]int, 0)
	track := make([]int, 0)
	visited := make(map[int]bool)

	var helper func()
	helper = func() {
		if len(track) == n {
			res = append(res, append([]int{}, track...))
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && visited[i-1] {
				continue
			}
			track = append(track, nums[i])
			visited[i] = true
			helper()
			track = track[:len(track)-1]
			visited[i] = false
		}
	}
	helper()

	return res
}

//39. 组合总和
//https://leetcode.cn/problems/combination-sum/description/
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	res := make([][]int, 0, 150)
	track := make([]int, 0)
	sum := 0
	var helper func(start int)
	helper = func(start int) {
		if sum == target {
			res = append(res, append([]int{}, track...))
			return
		}
		for i := start; i < n; i++ {
			if sum > target {
				break
			}
			track = append(track, candidates[i])
			sum += candidates[i]
			helper(i)
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}
	helper(0)
	return res
}

func Test_7(t *testing.T) {
	//fmt.Println(subsets([]int{1, 2, 3}))
	//fmt.Println(permute0([]int{1, 2, 3}))
	//fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	//fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	//fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
