package easy_test

import "testing"

func Test_hammingWeight2(t *testing.T) {
	println(hammingWeight2(7))
}

func hammingWeight2(num uint32) int {
	ans := 0
	for num != 0 {
		num = num & (num - 1)
		ans++
	}
	return ans
}
