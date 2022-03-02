package hard_test

import (
	"fmt"
	"testing"
)

func Test_checkRecord(t *testing.T) {
	fmt.Println(checkRecord(2))
}

var MOD = 1000000007

func checkRecord(n int) int {
	//P
	f := make([][][]int, n+1)
	for i := range f {
		//A
		f[i] = make([][]int, 2)
		for j := range f[i] {
			//L
			f[i][j] = make([]int, 3)
		}
	}
	f[0][0][0] = 1
	for i := 1; i < n+1; i++ {
		//P
		f[i][0][0] = (f[i-1][0][0] + f[i-1][0][1] + f[i-1][0][2]) % MOD
		f[i][1][0] = (f[i-1][1][0] + f[i-1][1][1] + f[i-1][1][2]) % MOD
		//A
		f[i][1][0] = (f[i][1][0] + f[i-1][0][0] + f[i-1][0][1] + f[i-1][0][2]) % MOD
		//L
		f[i][0][1] = f[i-1][0][0]
		f[i][0][2] = f[i-1][0][1]
		f[i][1][1] = f[i-1][1][0]
		f[i][1][2] = f[i-1][1][1]
	}

	ans := 0
	for i := range f[n] {
		for j := range f[n][i] {
			ans = (ans + f[n][i][j]) % MOD
		}
	}
	return ans
}

//DFS
func checkRecord0(n int) int {
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, 2)
		for j := range memo[i] {
			memo[i][j] = make([]int, 3)
		}
	}
	return checkRecordDFS(0, n, 0, 0, memo)
}

func checkRecordDFS(index, n, ACount, LCount int, memo [][][]int) int {
	chars := "ALP"
	if index >= n {
		return 1
	}
	if memo[index][ACount][LCount] != 0 {
		return memo[index][ACount][LCount]
	}
	var ans int
	for _, i := range chars {
		if i == 'A' && ACount < 1 {
			ans = (ans + checkRecordDFS(index+1, n, ACount+1, 0, memo)) % MOD
		} else if i == 'L' && LCount < 2 {
			ans = (ans + checkRecordDFS(index+1, n, ACount, LCount+1, memo)) % MOD
		} else if i == 'P' {
			ans = (ans + checkRecordDFS(index+1, n, ACount, 0, memo)) % MOD
		}
		memo[index][ACount][LCount] = ans
	}
	return ans
}
