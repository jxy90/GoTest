package Chapter1_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func minDistance(i, j int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if _s1[i] == _s2[j] {
		return minDistance(i-1, j-1)
	} else {
		return CommonUtil.Min(minDistance(i-1, j-1)+1,
			minDistance(i-1, j)+1,
			minDistance(i, j-1)+1)
	}
}

var _s1, _s2 string

func Test_minDistance(t *testing.T) {
	_s1 = "horse"
	_s2 = "ros"
	data := minDistance(len(_s1)-1, len(_s2)-1)
	println(data)
}
