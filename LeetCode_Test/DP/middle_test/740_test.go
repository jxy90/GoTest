package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"sort"
	"testing"
)

func Test_deleteAndEarn(t *testing.T) {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}

func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	counts := make([]int, 10001)
	max := 0
	for _, num := range nums {
		counts[num]++
		max = CommonUtil.Max(max, num)
	}
	f := make([][]int, max+1)
	for i := range f {
		f[i] = make([]int, 2)
	}

	for i := 1; i <= max; i++ {
		//但前值不被选择
		f[i][0] = CommonUtil.Max(f[i-1][0], f[i-1][1])
		//当前值被选择
		f[i][1] = f[i-1][0] + i*counts[i]
	}
	return CommonUtil.Max(f[max][0], f[max][1])
}
