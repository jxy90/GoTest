package middle_test

import (
	"math"
	"testing"
)

func Test_judgeSquareSum(t *testing.T) {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(judgeSquareSum(4))
	fmt.Println(judgeSquareSum(2))
}

func judgeSquareSum(c int) bool {
	temp := math.Sqrt(float64(c))
	left := 0
	right := int(math.Floor(temp))
	for left <= right {
		//mid := left + (right - left)
		if left*left+right*right == c {
			return true
		} else if left*left+right*right > c {
			right -= 1
		} else {
			left -= 1
		}
	}
	return false
}
