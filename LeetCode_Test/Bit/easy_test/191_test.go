package easy_test

import "testing"

func Test_hammingWeight(t *testing.T) {
	println(hammingWeight(1))
	println(hammingWeight(16))
	println(hammingWeight(218))
}
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
