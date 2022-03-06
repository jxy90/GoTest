package middle_test

import (
	"fmt"
	"testing"
)

func Test_2100(t *testing.T) {
	fmt.Println(goodDaysToRobBank([]int{5, 3, 3, 3, 5, 6, 2}, 2))
}

func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	g := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] == security[i-1] {
			continue
		} else if security[i] > security[i-1] {
			g[i] = 1
		} else {
			g[i] = -1
		}
	}
	a, b := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		if g[i-1] == 1 {
			a[i] = a[i-1] + 1
		}
		if g[i-1] == -1 {
			b[i] = b[i-1] + 1
		}
	}
	for i := time; i < n-time; i++ {
		c1 := a[i+1] - a[i+1-time]
		c2 := b[i+1+time] - b[i+1]
		if c1 == 0 && c2 == 0 {
			ans = append(ans, i)
		}
	}
	return
}
func goodDaysToRobBank0(security []int, time int) (ans []int) {
	n := len(security)
	for i := time; i < n-time; i++ {
		flag := true
		for j := i - time; flag && j < i; j++ {
			if security[j] < security[j+1] {
				flag = false
			}
		}
		for j := time; flag && j < i+time; j++ {
			if security[j] > security[j+1] {
				flag = false
			}
		}
		if flag {
			ans = append(ans, i)
		}
	}
	return
}
