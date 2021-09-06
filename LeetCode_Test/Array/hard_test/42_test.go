package hard_test

import (
	"fmt"
	"testing"
)

func Test_trip(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	ans := 0
	maxLeft := height[left]
	maxRight := height[right]
	for left < right {
		if height[left] > height[right] {
			right--
			if height[right] < min(maxLeft, maxRight) {
				ans += min(maxLeft, maxRight) - height[right]
			}
			maxRight = max(maxRight, height[right])
		} else {
			left++
			if height[left] < min(maxLeft, maxRight) {
				ans += min(maxLeft, maxRight) - height[left]
			}
			maxLeft = max(maxLeft, height[left])
		}

	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
