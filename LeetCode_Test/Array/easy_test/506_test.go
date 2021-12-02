package easy_test

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
}
func findRelativeRanks(score []int) []string {
	ssp := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	clone := make([]int, 0)
	for _, i := range score {
		clone = append(clone, i)
	}
	cache := map[int]string{}
	sort.Slice(clone, func(i, j int) bool {
		return clone[i] > clone[j]
	})
	for i, v := range clone {
		if i < 3 {
			cache[v] = ssp[i]
		} else {
			cache[v] = strconv.Itoa(i + 1)
		}
	}
	ans := make([]string, 0)
	for _, v := range score {
		ans = append(ans, cache[v])
	}
	return ans
}
