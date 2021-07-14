package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"sort"
	"testing"
)

func Test_minAbsoluteSumDiff(t *testing.T) {
	//[1,10,4,4,2,7]
	//[9,3,5,1,7,4]
	println(minAbsoluteSumDiff([]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}))
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	//这种方式不可用,是引用,并非复制
	//rec := sort.IntSlice(nums1)
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()
	sum, max := 0, 0
	//max = |num1[i]-nums2[i]|-|nums1[j]-nums2[i]|
	for i, v := range nums2 {
		a := CommonUtil.Abs(nums1[i] - v)
		sum += a
		//找到最接近v的值,保证max的至最大
		j := rec.Search(v)
		if j < n {
			max = CommonUtil.Max(max, a-CommonUtil.Abs(rec[j]-v))
		}
		if j > 0 {
			max = CommonUtil.Max(max, a-CommonUtil.Abs(rec[j-1]-v))
		}
	}
	return (sum - max) % (1e9 + 7)
}
