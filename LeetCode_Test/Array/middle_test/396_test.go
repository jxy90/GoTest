package middle_test

import (
	"testing"
)

func Test_396(t *testing.T) {
	t.Log(maxRotateFunction([]int{4, 3, 2, 6}))
	t.Log(maxRotateFunction([]int{100}))
	t.Log(maxRotateFunction([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

//f0=0*nums[0]+1*nums[1]+...+(n-2)*nums[n-2]+(n-1)*nums[n-1]
//f1=1*nums[0]+2*nums[1]+...+(n-1)*nums[n-2]+0*nums[n-1]=f0+sum(nums)-n*nums[n-1]
func maxRotateFunction0(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	sumNums := 0
	for i, num := range nums {
		f[0] += num * i
		sumNums += num
	}
	max := f[0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + sumNums - n*nums[n-i]
		if f[i] > max {
			max = f[i]
		}
	}
	return max
}

func maxRotateFunction(nums []int) int {
	n := len(nums)
	f := 0
	sumNums := 0
	for i, num := range nums {
		f += num * i
		sumNums += num
	}
	max := f
	for i := 1; i < n; i++ {
		f = f + sumNums - n*nums[n-i]
		if f > max {
			max = f
		}
	}
	return max
}
