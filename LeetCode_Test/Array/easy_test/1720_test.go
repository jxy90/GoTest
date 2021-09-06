package easy_test

import (
	"fmt"
	"testing"
)

func Test_decode(t *testing.T) {
	fmt.Println(decode([]int{1, 2, 3}, 1))
}

func decode(encoded []int, first int) []int {
	ans := make([]int, 0)
	ans = append(ans, first)
	for i := range encoded {
		first = first ^ encoded[i]
		ans = append(ans, first)
	}
	return ans
}
