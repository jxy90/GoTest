package easy_test

import (
	"fmt"
	"testing"
)

func Test_maxNumberOfBalloons(t *testing.T) {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
}

//balloon
func maxNumberOfBalloons(text string) int {
	cache := map[byte]int{}
	for i := range text {
		b := text[i]
		if b == 'b' || b == 'a' || b == 'l' || b == 'o' || b == 'n' {
			cache[b]++
		}
	}
	cache['l'] /= 2
	cache['o'] /= 2
	ans := cache['b']
	for _, i := range cache {
		if ans > i {
			ans = i
		}
	}
	return ans
}
