package hard_test

import (
	"sort"
	"testing"
)

func Test_899(t *testing.T) {

}

func orderlyQueue(s string, k int) string {
	if k == 1 {
		n := len(s)
		ans := s
		for i := 0; i < n; i++ {
			s = s[1:] + s[:1]
			if ans > s {
				ans = s
			}
		}
		return ans
	}
	temp := []byte(s)
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	return string(temp)
}
