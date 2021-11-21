package middle_test

import (
	"fmt"
	"testing"
)

func Test_integerReplacement(t *testing.T) {
	fmt.Println(integerReplacement(8))
}

func integerReplacement0(n int) (ans int) {
	for n != 1 {
		if n%2 == 1 {
			if n != 3 && (n>>1)&1 == 1 {
				n++
			} else {
				n--
			}
		} else {
			n >>= 1
		}
		ans++
	}
	return
}
func integerReplacement(n int) (ans int) {
	queue := make([][]int, 0)
	queue = append(queue, []int{n, 0})
	visited := map[int]bool{}
	visited[n] = true
	for len(queue) > 0 {
		curVal := queue[0][0]
		curStep := queue[0][1]
		if curVal == 1 {
			return curStep
		}
		queue = queue[1:]
		newVal := 0
		if curVal%2 == 0 {
			newVal = curVal / 2
			if !visited[newVal] {
				queue = append(queue, []int{newVal, curStep + 1})
			}
		} else {
			newVal = curVal + 1
			if !visited[newVal] {
				queue = append(queue, []int{newVal, curStep + 1})
			}
			newVal = curVal - 1
			if !visited[newVal] {
				queue = append(queue, []int{newVal, curStep + 1})
			}
		}
	}
	return -1
}
