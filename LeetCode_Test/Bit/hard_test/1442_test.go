package hard_test

import (
	"fmt"
	"testing"
)

func Test_countTriplets(t *testing.T) {
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))
	fmt.Println(countTriplets([]int{1, 1, 1, 1, 1}))
	fmt.Println(countTriplets([]int{2, 3}))
	fmt.Println(countTriplets([]int{1, 3, 5, 7, 9}))
	fmt.Println(countTriplets([]int{7, 11, 12, 9, 5, 2, 7, 17, 22}))
}

func countTriplets(arr []int) int {
	n := len(arr)
	f := make([]int, n)
	f[0] = arr[0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] ^ arr[i]
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for k := i + 1; k < n; k++ {
			if f[k]^f[i]^arr[i] == 0 {
				ans += k - i
			}
		}
	}
	return ans
}
func countTriplets1(arr []int) int {
	n := len(arr)
	f := make([]int, n)
	f[0] = arr[0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] ^ arr[i]
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				if f[k]^f[i]^arr[i] == 0 {
					ans++
				}
			}
		}
	}
	return ans
}

func countTriplets0(arr []int) int {
	n := len(arr)
	f := make([]int, n)
	f[0] = arr[0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] ^ arr[i]
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				a := f[j] ^ f[i] ^ arr[j] ^ arr[i]
				b := f[k] ^ f[j-1]
				if a == b {
					ans++
				}
			}
		}
	}
	return ans
}
