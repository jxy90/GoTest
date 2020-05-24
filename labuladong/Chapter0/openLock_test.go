package Chapter0_test

import "testing"

func openLock(deadends []string, target string) int {
	queue := []string{}
	queue = append(queue, "0000")
	depth := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i] == target {
				return depth
			} else if openLockContains(deadends, queue[i]) {
				continue
			}
			for j := 0; j < len(queue[i]); j++ {
				queue = append(queue, openLockUp(queue[i], j))
				queue = append(queue, openLockDown(queue[i], j))
			}
		}
		queue = queue[length:]
		depth++
	}
	return depth
}

func openLockContains(deadends []string, target string) bool {
	for _, b := range deadends {
		if b == target {
			return true
		}
	}
	return false
}

func openLockUp(str string, index int) string {
	ab := make([]int32, 0)
	for _, v := range str {
		ab = append(ab, v)
	}
	if ab[index] == '9' {
		ab[index] = '0'
	} else {
		ab[index] += 1
	}
	return string(ab)
}

func openLockDown(str string, index int) string {
	ab := make([]int32, 0)
	for _, v := range str {
		ab = append(ab, v)
	}
	if ab[index] == '0' {
		ab[index] = '9'
	} else {
		ab[index] -= 1
	}
	return string(ab)
}

func Test_openLock(t *testing.T) {
	println(openLockUp("0000", 2))
	println(openLockDown("0000", 2))
}
