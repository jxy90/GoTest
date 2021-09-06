package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_numRescueBoats(t *testing.T) {
	fmt.Println(numRescueBoats([]int{3, 1, 7}, 7))
	//fmt.Println(numRescueBoats([]int{3, 2, 3, 2, 2}, 6))
	//fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	n := len(people)
	left, right := 0, n-1
	count := 0
	weight := 0
	preCount := 0
	for left <= right {
		for left <= right {
			if preCount == 2 {
				break
			}
			if weight+people[right] <= limit {
				weight += people[right]
				right--
				preCount++
			} else if weight+people[left] <= limit {
				weight += people[left]
				left++
				preCount++
			} else {
				break
			}
		}
		count++
		weight = 0
		preCount = 0
	}
	return count
}
