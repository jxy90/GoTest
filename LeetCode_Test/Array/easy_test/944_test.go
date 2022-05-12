package easy_test

import (
	"fmt"
	"testing"
)

func Test_944(t *testing.T) {
	fmt.Println(minDeletionSize([]string{"abc", "bce", "cae"}))
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
}

func minDeletionSize(strs []string) int {
	ans := 0
	for j := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] > strs[i][j] {
				ans++
				break
			}
		}
	}
	return ans
}
