package Array

import (
	"testing"
)

//按权重随机选择

//type Solution struct {
//	preSum []int
//}
//
//func Constructor(w []int) Solution {
//	preSum := make([]int, len(w)+1)
//	for i := 1; i <= len(w); i++ {
//		preSum[i] = preSum[i-1] + w[i-1]
//	}
//	return Solution{preSum}
//}
//
//func (this *Solution) PickIndex() int {
//	target := rand.Intn(this.preSum[len(this.preSum)-1]) + 1
//	left, right := 0, len(this.preSum)
//	for left <= right {
//		mid := left + (right-left)>>1
//		val := this.preSum[mid]
//		if val == target {
//			right = mid
//		} else if val > target {
//			right = mid - 1
//		} else if val < target {
//			left = mid + 1
//		}
//	}
//	return left
//}

func Test_6(t *testing.T) {
}
