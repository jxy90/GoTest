package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_escapeGhosts(t *testing.T) {
	println(escapeGhosts([][]int{{1, 0}, {0, 3}}, []int{0, 1}))
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	distance := mhtDistance([]int{0, 0}, target)
	for _, ghost := range ghosts {
		if mhtDistance(ghost, target) <= distance {
			return false
		}
	}
	return true
}

func mhtDistance(s, t []int) int {
	return CommonUtil.Abs(s[0]-t[0]) + CommonUtil.Abs(s[1]-t[1])
}
