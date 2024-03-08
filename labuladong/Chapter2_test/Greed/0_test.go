package Greed

import (
	"fmt"
	"math"
	"testing"
)

func min(nums ...int) int {
	min := math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(nums ...int) int {
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

//134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	diff := make([]int, len(gas))
	index := 0
	for i := 0; i < len(gas); i++ {
		diff[i] = gas[i] - cost[i]
		if diff[i] > diff[index] {
			index = i
		}
	}
	newDiff := append(diff[index:], diff[:index]...)
	sum := 0
	for _, i := range newDiff {
		sum += i
		if sum < 0 {
			return -1
		}
	}
	return index
}
func Test_0(t *testing.T) {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
