package easy_test

import (
	"testing"
)

func Test_guessNumber(t *testing.T) {
	data := guessNumber(3)
	fmt.Println(data)
}

func guess(num int) int {
	return 1
}
func guessNumber(n int) int {
	m := 0
	for m <= n {
		temp := (m + n) / 2
		if guess(temp) == 0 {
			return temp
		} else if guess(temp) < 0 {
			n = temp - 1
		} else {
			m = temp + 1
		}
	}
	return m
}
