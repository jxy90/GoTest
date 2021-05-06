package easy_test

import "testing"

func Test_decode(t *testing.T) {
	println(decode([]int{1, 2, 3}, 1))
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
