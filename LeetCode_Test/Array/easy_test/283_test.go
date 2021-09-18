package easy_test

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	//nums := []int{0, 1, 0, 3, 12}
	nums := []int{0, 0, 1}
	moveZeroes3(nums)
	for i := range nums {
		fmt.Println(nums[i])
	}
}
func moveZeroes(nums []int) {
	count := 0 //0的数量
	index := 0 //0的起始位置
	for i := range nums {
		//找到第一个0的位置
		if count == 0 && nums[i] == 0 {
			index = i
			count++
			continue
		}
		//判断当前位是否为0
		if nums[i] == 0 {
			//是
			count++
			continue
		} else {
			//否
			for j := 0; j < count; j++ {
				nums[i+j], nums[index+j] = nums[index+j], nums[i+j]
			}
			index++
		}

		if index+count == len(nums) {
			break
		}
	}
}

func moveZeroes3(nums []int) {
	n, left, right := len(nums), 0, 0
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

}
