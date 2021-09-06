package easy_test

import (
	"fmt"
	"testing"
)

func Test_xorOperation(t *testing.T) {
	fmt.Println(xorOperation(5, 0))
}
func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= start + i*2
	}
	return ans
}
