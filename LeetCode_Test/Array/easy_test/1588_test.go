package easy_test

import (
	"fmt"
	"testing"
)

func Test_sumOddLengthSubarrays(t *testing.T) {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}
func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + arr[i]
	}
	ans := 0
	for length := 1; length <= n; length += 2 {
		for i := 0; i+length <= n; i++ {
			ans += sum[i+length] - sum[i]
		}
	}
	return ans
}
