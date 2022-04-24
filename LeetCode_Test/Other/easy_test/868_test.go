package easy_test

import (
	"log"
	"testing"
)

func Test_868(t *testing.T) {
	log.Println(binaryGap(22))
	log.Println(binaryGap(8))
	log.Println(binaryGap(5))
}

func binaryGap(n int) (ans int) {
	first, second := 100, 100
	//替换第一位和第二位1的index，并计算ans
	var helper = func(index int) {
		second = first
		first = index
		if first-second > ans {
			ans = first - second
		}
	}
	index := 0
	for n > 0 {
		//每次取第一位是否为1
		cur := n & 1
		if cur == 1 {
			helper(index)
		}
		n >>= 1
		index++
	}
	return ans
}
