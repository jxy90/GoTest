package easy_test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	fmt.Println(fizzBuzz(123))
}

func fizzBuzz(n int) (ans []string) {
	for i := 1; i <= n; i++ {
		temp := ""
		if i%3 == 0 {
			temp += "Fizz"
		}
		if i%5 == 0 {
			temp += "Buzz"
		}
		if temp != "" {
			ans = append(ans, temp)
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return
}
