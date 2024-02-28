package nSum

import (
	"fmt"
	"sort"
	"testing"
)

//twoSum 问题
//https://labuladong.online/algo/practice-in-action/nsum/
//https://cloud.tencent.com/developer/article/1880894

func Test_0(t *testing.T) {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))

	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	fmt.Println(nSum(nums, 3, 0, 0))
}

//2、两数之和
//https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	ans := make([]int, 0, 2)
	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			ans = append(ans, left)
			ans = append(ans, right)
			return ans
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return ans
}

//nums 中可能有多对儿元素之和都等于 target，请你的算法返回所有和为 target 的元素对儿，其中不能出现重复。
func twoSum2(numbers []int, target int) [][]int {
	left, right := 0, len(numbers)-1
	ans := make([][]int, 0)
	for left < right {
		leftVal := numbers[left]
		rightVal := numbers[right]
		sum := leftVal + rightVal
		if sum == target {
			ans = append(ans, []int{numbers[left], numbers[right]})
			for left < right && left+1 < len(numbers) && numbers[left] == leftVal {
				left++
			}
			for left < right && right-1 > 0 && numbers[right] == rightVal {
				right--
			}
		} else if sum < target {
			for left < right && left+1 < len(numbers) && numbers[left] == leftVal {
				left++
			}
		} else if sum > target {
			for left < right && right-1 > 0 && numbers[right] == rightVal {
				right--
			}
		}
	}
	return ans
}

func nSum(numbers []int, n, start, target int) [][]int {
	ans := make([][]int, 0)
	if n < 2 || n > len(numbers) {
		return ans
	}
	if n == 2 {
		left, right := start, len(numbers)-1
		for left < right {
			leftVal, rightVal := numbers[left], numbers[right]
			sum := leftVal + rightVal
			if sum == target {
				ans = append(ans, []int{numbers[left], numbers[right]})
				for left < right && numbers[left] == leftVal {
					left++
				}
				for left < right && numbers[right] == rightVal {
					right--
				}
			} else if sum < target {
				for left < right && numbers[left] == leftVal {
					left++
				}
			} else if sum > target {
				for left < right && numbers[right] == rightVal {
					right--
				}
			}
		}
	} else {
		for i := start; i < len(numbers); i++ {
			tempAns := nSum(numbers, n-1, i+1, target-numbers[i])
			for j := range tempAns {
				tempAns[j] = append(tempAns[j], numbers[i])
				ans = append(ans, tempAns[j])
			}
			for i < len(numbers)-1 && numbers[i] == numbers[i+1] {
				i++
			}
		}
	}
	return ans
}
