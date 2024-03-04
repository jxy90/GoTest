package Tree

import (
	"fmt"
	"testing"
)

//数组中的第K个最大元素
//https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
func findKthLargest(nums []int, k int) int {
	findKthLargestHelper(nums, 0, len(nums)-1, k)
	return nums[len(nums)-k]
}

func findKthLargestHelper(nums []int, start, end, k int) {
	if start < end {
		left, right := start, end
		mid := nums[(left+right)/2]
		for left <= right {
			for nums[left] < mid {
				left++
			}
			for nums[right] > mid {
				right--
			}
			if left <= right {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right--
			}
		}
		if len(nums)-k < right {
			findKthLargestHelper(nums, start, right, k)
		} else {
			findKthLargestHelper(nums, left, end, k)
		}
	}
}

//排序数组
//https://leetcode.cn/problems/sort-an-array/description/
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		left, right := start, end
		mid := nums[left+(right-left)>>1]
		for left <= right {
			for nums[left] < mid {
				left++
			}
			for nums[right] > mid {
				right--
			}
			if left <= right {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right--
			}
		}
		if right > start {
			quickSort(nums, start, right)
		}
		if left < end {
			quickSort(nums, left, end)
		}
	}
}

//二叉搜索树的最近公共祖先
//https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	min, max := 0, 0
	if p.Val < q.Val {
		min = p.Val
		max = q.Val
	} else {
		min = q.Val
		max = p.Val
	}
	target := root
	for {
		if root.Val >= min && root.Val <= max {
			target = root
			break
		} else if root.Val < min {
			root = root.Right
		} else if root.Val > max {
			root = root.Left
		}
	}
	return target
}

//https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return root.Left
	} else {
		return root.Right
	}
}

func Test_9(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	//fmt.Println(findKthLargest(nums, 2))
	//fmt.Println(findKthLargest([]int{3, 1, 2, 4}, 2))
	fmt.Println(sortArray(nums))
}
