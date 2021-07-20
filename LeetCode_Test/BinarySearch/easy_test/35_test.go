package easy_test

import (
	"testing"
)

func Test_searchInsert(t *testing.T) {
	println(searchInsert([]int{1, 3, 5, 6}, 5))
	println(searchInsert([]int{1, 3, 5, 6}, 2))
	println(searchInsert([]int{1, 3, 5, 6}, 7))
	println(searchInsert([]int{1, 3, 5, 6}, 0))
}

//把问题想象成找到 >=target的位置
func searchInsert(nums []int, target int) int {
	//确定左右端点
	l, r := 0, len(nums)-1
	//如果nums最大值小于target返回数组最右的位置
	if nums[r] < target {
		return r + 1
	}
	for l < r {
		//取中间位置
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			//此位置可能满足条件,所以不能-1
			r = mid
		} else {
			//此位置不满子条件.所以可以+1
			l = mid + 1
		}
	}
	//这时l==r即为结果
	return l
}
