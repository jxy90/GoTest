package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	tails := make([]int, n)
	ans := 0
	for _, num := range nums {
		i, j := 0, ans
		for i < j {
			mid := i + (j-i)/2
			if tails[mid] < num {
				i = mid + 1
			} else {
				j = mid
			}
		}
		tails[i] = num
		if ans == j {
			ans++
		}
	}
	return ans
}

func lengthOfLIS0(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		f[i] = 1
		for j := 1; j < i; j++ {
			if nums[i-1] >= nums[j-1] {
				f[i] = CommonUtil.Max(f[i], f[j]+1)
			}
		}
	}
	ans := 0
	for i := range f {
		ans = CommonUtil.Max(ans, f[i])
	}
	return ans
}
