package middle_test

import (
	"testing"
)

func Test_735(t *testing.T) {
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	for _, asteroid := range asteroids {
		alive := true
		for alive && asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			alive = -asteroid > stack[len(stack)-1]
			if -asteroid >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
		if alive {
			stack = append(stack, asteroid)
		}
	}
	return stack
}
