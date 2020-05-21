package hard_test

import "testing"

func minDistance(word1 string, word2 string) int {
	_word1, _word2 = word1, word2
	return minDistanceDP(len(word1)-1, len(word2)-1)
}

var _word1, _word2 string

func minDistanceDP(i, j int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if _word1[i] == _word2[j] {
		return minDistanceDP(i-1, j-1)
	} else {
		return min(
			minDistanceDP(i, j-1)+1,   //insert
			minDistanceDP(i-1, j)+1,   //delete
			minDistanceDP(i-1, j-1)+1, //replace
		)
	}
}

func min(args ...int) int {
	min := args[0]
	for k := range args {
		if min > args[k] {
			min = args[k]
		}
	}
	return min
}


func Test_minDistance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	println(minDistance(word1,word2))
	word1 = "intention"
	word2 = "execution"
	println(minDistance(word1,word2))
}
