package easy_test

import (
	"fmt"
	"testing"
)

func Test_readBinaryWatchHelper(t *testing.T) {
	fmt.Println(readBinaryWatch(4))
}

//var readBinaryWatchOptions = []int{1, 2, 4, 8, 16, 32, 60, 120, 240, 480}

func readBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if readBinaryWatchCount(i)+readBinaryWatchCount(j) == turnedOn {
				jj := ""
				if j < 10 {
					jj = fmt.Sprintf("0%d", j)
				} else {
					jj = fmt.Sprintf("%d", j)
				}
				temp := fmt.Sprintf("%d:"+jj, i)
				ans = append(ans, temp)
			}
		}
	}
	return ans
}

func readBinaryWatchCount(n int) int {
	res := 0
	for n != 0 {
		n = n & (n - 1)
		res++
	}
	return res
}
