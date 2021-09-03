package hard_test

import (
	"testing"
)

var memo = make(map[string]bool)

func isScramble(s1 string, s2 string) bool {
	key := s1 + ":" + s2
	if v, ok := memo[key]; ok {
		return v
	}
	if len(s1) != len(s2) {
		memo[key] = false
		return false
	}
	if s1 == s2 {
		memo[key] = true
		return true
	}
	n := len(s1)
	for i := 1; i < n; i++ {
		res1 := isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])
		res2 := isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i])
		if res1 || res2 {
			memo[key] = true
			return true
		}
	}
	memo[key] = false
	return false
}

//func isScramble(s1 string, s2 string) bool {
//	n := len(s1)
//	dp := make([][][]int, n)
//	for i := range dp {
//		dp[i] = make([][]int, n)
//		for j := range dp[i] {
//			dp[i][j] = make([]int, n+1)
//			for k := range dp[i][j] {
//				dp[i][j][k] = -1
//			}
//		}
//	}
//	for i := 1; i < n; i++ {
//		// 不交换的情况
//		if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
//			return 1
//		}
//		// 交换的情况
//		if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
//			return 1
//		}
//	}
//}

func isScrambleV(s1, s2 string, i, j, k int, dp [][][]int) int {
	if dp[i][j][k] == -1 {
		if hashCode(s1[i:i+k]) == hashCode(s2[j:j+k]) {
			dp[i][j][k] = 1
		} else {
			dp[i][j][k] = 0
		}
	}
	return dp[i][j][k]
}

func isScrambleHelper(s1 string, s2 string) bool {
	key := s1 + s2
	if value, ok := memo[key]; ok {
		return value
	}
	n := len(s1)
	if n <= 1 {
		return isScrambleValidate(s1, s2)
	}
	ans := false
	for i := 1; i < n; i++ {
		ff := isScrambleValidate(s1[:i], s2[:i]) && isScrambleValidate(s1[i:], s2[i:])
		if ff {
			ff = isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])
		}
		fe := isScrambleValidate(s1[:i], s2[n-i:]) && isScrambleValidate(s1[i:], s2[:n-i])
		if fe {
			fe = isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i])
		}
		ans = ans || ff || fe
	}
	return ans
}

func hashCode(s string) int {
	x := 0
	for i := range s {
		x += 1 << (s[i] - 'a')
	}
	return x
}

func isScrambleValidate(s1 string, s2 string) bool {
	//for i := range s1 {
	//	if !strings.Contains(s2, string(s1[i])) {
	//		return false
	//	}
	//}
	//return true

	return true
}

func Test_isScramble(t *testing.T) {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abcde", "caebd"))
	fmt.Println(isScramble("a", "a"))
	fmt.Println(isScramble("aa", "aa"))
	fmt.Println(isScramble("aa", "ab"))
	fmt.Println(isScramble("ba", "ab"))
	fmt.Println(isScramble("abcdd", "dbdac"))
}
