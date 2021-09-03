package easy_test

import (
	"testing"
)

func Test_getMaximumGenerated(t *testing.T) {
	fmt.Println(getMaximumGenerated(4))
}

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	ans := 0
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}
