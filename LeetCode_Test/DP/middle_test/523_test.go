package middle_test

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
	fmt.Println(checkSubarraySum([]int{1, 0}, 2))
	fmt.Println(checkSubarraySum([]int{5, 0, 0, 0}, 3))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
}

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	f := make(map[int]int)
	f[0] = -1
	rem := 0
	for i := 0; i < n; i++ {
		rem += nums[i]
		if index, ok := f[rem%k]; ok {
			if i-index > 1 {
				return true
			}
		} else {
			f[rem%k] = i
		}
	}
	return false
}
func checkSubarraySum0(nums []int, k int) bool {
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] + nums[i-1]
	}
	m := len(f)
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if i == 0 && j == m-1 {
				continue
			}
			if (f[j]-f[i])%k == 0 {
				return true
			}
		}
	}
	return false
}
