package hard_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"sort"
	"testing"
)

func Test_minimumTimeRequired(t *testing.T) {
	//println(minimumTimeRequired([]int{3, 2, 3}, 3))
	println(minimumTimeRequired([]int{1, 2, 4, 7, 8}, 2))
	println(minimumTimeRequired([]int{5, 5, 4, 4, 4}, 2))
}

var ans int

func minimumTimeRequired(jobs []int, k int) int {
	ans = math.MaxInt32
	minimumTimeRequiredDFS(jobs, make([]int, k), 0, 0, 0)
	return ans
}

func minimumTimeRequiredDFS(jobs, sum []int, index, max, used int) {
	if max > ans {
		return
	}
	if index == len(jobs) {
		ans = CommonUtil.Min(max, ans)
		return
	}
	if used < len(sum) {
		sum[used] = jobs[index]
		minimumTimeRequiredDFS(jobs, sum, index+1, CommonUtil.Max(max, jobs[index]), used+1)
		sum[used] = 0
	}
	for i := 0; i < used; i++ {
		sum[i] += jobs[index]
		minimumTimeRequiredDFS(jobs, sum, index+1, CommonUtil.Max(max, sum[i]), used)
		sum[i] -= jobs[index]
	}
}

func minimumTimeRequired0(jobs []int, k int) int {
	sort.Ints(jobs)
	n := len(jobs)
	sum := make([]int, k)
	for i := n - 1; i >= 0; i-- {
		index := minIndex(sum)
		sum[index] += jobs[i]
	}
	return sum[maxIndex(sum)]
}

func minIndex(sum []int) int {
	index := 0
	for i := 1; i < len(sum); i++ {
		if sum[i] < sum[index] {
			index = i
		}
	}
	return index
}
func maxIndex(sum []int) int {
	index := 0
	for i := 1; i < len(sum); i++ {
		if sum[i] > sum[index] {
			index = i
		}
	}
	return index
}
