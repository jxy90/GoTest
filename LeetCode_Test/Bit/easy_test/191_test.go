package easy_test

import (
	"fmt"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	fmt.Println(hammingWeight(1))
	fmt.Println(hammingWeight(16))
	fmt.Println(hammingWeight(218))
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
