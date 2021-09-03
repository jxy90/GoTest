package middle_test

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Println(numberOfArithmeticSlices([]int{1}))
}

//子数组,找到连续等差最长的一个子数组
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := 0
	length := 0
	diff := nums[0] - nums[1]
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == diff {
			//求差相,组数长度+1
			length++
		} else {
			//不相等,重新记录差值,重置记录长度
			diff = nums[i-1] - nums[i]
			length = 0
		}
		ans += length
	}
	return ans
}
