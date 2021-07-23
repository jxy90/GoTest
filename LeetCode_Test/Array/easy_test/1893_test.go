package easy_test

import (
	"sort"
	"testing"
)

func Test_isCovered(t *testing.T) {
	println(isCovered([][]int{{1, 2}, {3, 4}, {4, 5}}, 2, 5))
}

//差分
func isCovered(ranges [][]int, left int, right int) bool {
	diff := make([]int, 52)
	for _, ints := range ranges {
		diff[ints[0]]++
		diff[ints[1]+1]--
	}
	sum := make([]int, 52)
	for i := 1; i < len(diff); i++ {
		sum[i] = sum[i-1] + diff[i]
	}
	for i := left; i <= right; i++ {
		if sum[i] == 0 {
			return false
		}
	}
	return true
}

//0ms 2.5MB
//减少目标区间
func isCovered2(ranges [][]int, left int, right int) bool {
	//升序排序
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	for _, ints := range ranges {
		if ints[1] < left || ints[0] > right {
			//ints在目标的左侧或者右侧,没有交集
			continue
		} else if ints[0] <= left && ints[1] >= right {
			//ints完全覆盖目标区间
			return true
		} else if ints[0] <= left && ints[1] >= left {
			//ints部分覆盖目标区间,减少目标区间
			left = ints[1] + 1
		} else if ints[0] > left {
			//ints不能覆盖目标的左端,因为前面排序,失败
			return false
		}
	}
	return false
}

//0ms 2.6MB
//哈希
func isCovered0(ranges [][]int, left int, right int) bool {
	cache := map[int]bool{}
	for i := left; i <= right; i++ {
		cache[i] = false
	}
	for _, item := range ranges {
		for i := item[0]; i <= item[1]; i++ {
			if _, ok := cache[i]; ok {
				cache[i] = true
			}
		}
	}
	for _, val := range cache {
		if !val {
			return false
		}
	}
	return true
}
