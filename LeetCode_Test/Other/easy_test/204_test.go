package easy_test

import (
	"fmt"
	"testing"
)

func countPrimes(n int) int {
	ans := 0
	nums := make([]int, 0)
	if n == 0 || n == 1 {
		return ans
	}
	for i := 2; i < n; i++ {
		if countPrimesHelper(nums, i) {
			ans++
			nums = append(nums, i)
		}
	}
	for i := range nums {
		fmt.Println(nums[i])
	}
	return ans
}

func countPrimesHelper(nums []int, n int) bool {
	//if n != 2 && n%2 == 0 && n%3 == 0 {
	//	return false
	//}
	maxNum := 2
	for i := range nums {
		maxNum = nums[i]
		if n != maxNum && n%maxNum == 0 {
			return false
		}
	}
	//half := n / 2
	for i := maxNum; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Test_countPrimes(t *testing.T) {
	countPrimes(10)
	countPrimes2(10)
}

func countPrimes2(n int) int {
	isPrim := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrim[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrim[i] == true {
			for j := i * i; j < n; j = j + i {
				isPrim[j] = false
			}
		}
	}
	ans := 0
	for _, b := range isPrim {
		if b {
			ans++
		}
	}
	fmt.Println(ans)
	return ans
}
