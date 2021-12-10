package hard_test

import (
	"fmt"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	n := len(nums)
	right := 0
	//保存最大字的索引
	queue := make([]int, 0)
	for right < n {
		rightVal := nums[right]
		//如果当前值>=queue中的索引对应的值,把他们都弹出
		for len(queue) > 0 && rightVal >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		//加入当前值
		queue = append(queue, right)
		//判断queue首位的索引在不在窗口内,不在移除
		if right-k >= queue[0] {
			queue = queue[1:]
		}
		//从滑动窗口内取最大值
		if right-k+1 >= 0 {
			ans = append(ans, nums[queue[0]])
		}
		right++
	}
	return ans
}
