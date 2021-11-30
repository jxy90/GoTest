package middle_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Test_findNthDigit(t *testing.T) {
	fmt.Println(findNthDigit0(3))
	fmt.Println(findNthDigit0(11))
	fmt.Println(findNthDigit0(101))
}

func findNthDigit(n int) int {
	//求位数d
	//n等于当前为数的index
	d := 1
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}
	//
	index := n - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex-1)) % 10
}
func findNthDigit0(n int) int {
	left, right := 0, math.MaxInt16
	for left < right {
		midNum := left + (right-left)>>1
		start, end := findNthDigitHelper(midNum)
		if start <= n && n < end {
			//中!
			strN := strconv.Itoa(midNum)
			for i := 0; i < len(strN); i++ {
				if i+start == n {
					ans, _ := strconv.Atoi(string(strN[i]))
					return ans
				}
			}
		} else if n < start {
			right = midNum - 1
		} else {
			left = midNum
		}
	}
	return -1
}

func findNthDigitHelper(num int) (int, int) {
	//求位数
	cnt := 1
	tempNum := num
	base := 9
	start := 0
	for tempNum > 9 {
		start += base * cnt

		base = base * 10
		tempNum /= 10
		cnt++
	}

	chu := 1
	for i := 0; i < cnt-1; i++ {
		chu *= 10
	}
	start += (num-chu)*cnt + 1
	end := start + cnt
	//前闭后开
	return start, end
}
