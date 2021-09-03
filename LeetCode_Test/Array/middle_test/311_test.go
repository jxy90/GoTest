package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

func Test_nthSuperUglyNumber(t *testing.T) {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

func nthSuperUglyNumber(n int, primes []int) int {
	m := len(primes)
	p := make([]int, m)
	nums := make([]int, m)
	for i := range p {
		p[i] = 1
	}
	f := make([]int, n+1)
	f[1] = 1

	for i := 2; i <= n; i++ {
		min := math.MaxInt32
		for j := range nums {
			nums[j] = f[p[j]] * primes[j]
			min = CommonUtil.Min(min, nums[j])
		}
		f[i] = min
		for j := range nums {
			if f[i] == nums[j] {
				p[j]++
			}
		}
	}
	return f[n]
}

func nthSuperUglyNumber0(n int, primes []int) int {
	index := 1
	cnt := 0
	for {
		if isNthSuperUglyNumber(index, primes) {
			cnt++
		}
		if cnt == n {
			return index
		}
		index++
	}
	return -1
}

func isNthSuperUglyNumber(num int, primes []int) bool {
	if num <= 0 {
		return false
	}
	for _, prime := range primes {
		for num%prime == 0 {
			num /= prime
		}
	}
	return num == 1
}
