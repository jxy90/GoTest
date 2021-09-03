package Chapter1_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func minDistanceDp(i, j int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if _s1[i] == _s2[j] {
		memo[i][j] = minDistanceDp(i-1, j-1)
	} else {
		memo[i][j] = CommonUtil.Min(minDistanceDp(i-1, j-1)+1,
			minDistanceDp(i-1, j)+1,
			minDistanceDp(i, j-1)+1)
	}
	return memo[i][j]
}

func minDistance(s1, s2 string) int {
	return minDistanceDp(len(_s1)-1, len(_s2)-1)
}

var _s1, _s2 string
var memo [][]int

func Test_minDistance(t *testing.T) {
	_s1 = "horse"
	_s2 = "ros"
	memo = make([][]int, len(_s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(_s2)+1)
	}
	data := minDistance(_s1, _s2)
	fmt.Println(data)

	data = minDistanceDP(_s1, _s2)
	fmt.Println(data)
}

//DP
func minDistanceDP(word1, word2 string) int {
	n1, n2 := len(word1), len(word2)

	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = CommonUtil.Min(
					dp[i-1][j-1]+1,
					dp[i-1][j]+1,
					dp[i][j-1]+1)
			}
		}
	}

	return dp[n1][n2]
}
