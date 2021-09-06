package middle_test

import (
	"fmt"
	"testing"
)

func Test_countArrangement(t *testing.T) {
	//1,1
	//2,2
	//3,
	fmt.Println(countArrangement(1))
}

func countArrangement(n int) int {
	visited := make([]bool, n+1)
	match := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		match[i] = make([]int, 0)
		for j := 1; j < n+1; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}

	ans := 0
	var loop func(int)
	loop = func(index int) {
		if index > n {
			ans++
			return
		}
		for _, v := range match[index] {
			if visited[v] {
				continue
			}
			visited[v] = true
			loop(index + 1)
			visited[v] = false
		}
	}
	loop(1)
	return ans
}
