package middle_test

import (
	"fmt"
	"testing"
)

func Test_getHint(t *testing.T) {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
}

func getHint(secret string, guess string) string {
	x, y := 0, 0
	n := len(secret)
	cache := map[byte]int{}
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			x++
		}
		cache[secret[i]]++
	}
	for i := range guess {
		if cache[guess[i]] != 0 {
			y++
			cache[guess[i]]--
		}
	}
	return fmt.Sprintf("%vA%vB", x, y-x)
}
