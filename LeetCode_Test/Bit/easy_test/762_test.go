package easy_test

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	fmt.Println(countPrimeSetBits(6, 10))
}

func countPrimeSetBits(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		if isSmallPrime(getBitCount(i)) {
			ans++
		}
	}
	return ans
}

func isSmallPrime(x int) bool {
	return x == 2 || x == 3 || x == 5 || x == 7 || x == 11 || x == 13 || x == 17 || x == 19
}

func getBitCount(n int) int {
	res := 0
	for n > 0 {
		if n&1 == 1 {
			res++
		}
		n >>= 1
	}

	return res
}
