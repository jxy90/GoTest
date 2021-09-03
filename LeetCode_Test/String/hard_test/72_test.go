package hard_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func minDistance(word1 string, word2 string) int {
	_word1, _word2 = word1, word2
	length1, length2 := len(_word1)-1, len(_word2)-1

	_memo = make([][]int, len(_word1), len(_word1))
	//fmt.Println(len(_memo))
	//fmt.Println(len(word1))
	for i := 0; i < len(_word1); i++ {
		_memo[i] = make([]int, len(_word2), len(_word2))
	}

	return minDistanceDP(length1, length2)
}

var _word1, _word2 string
var _memo [][]int

func minDistanceDP(i, j int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	minDp := 0
	if _memo[i][j] != 0 {
		return _memo[i][j]
	}
	if _word1[i] == _word2[j] {
		minDp = minDistanceDP(i-1, j-1)
	} else {
		minDp = CommonUtil.Min(
			minDistanceDP(i, j-1)+1,   //insert
			minDistanceDP(i-1, j)+1,   //delete
			minDistanceDP(i-1, j-1)+1, //replace
		)
	}
	_memo[i][j] = minDp
	return minDp
}

func Test_minDistance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
	word1 = "intention"
	word2 = "execution"
	fmt.Println(minDistance(word1, word2))
}
