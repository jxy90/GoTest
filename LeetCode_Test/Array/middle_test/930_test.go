package middle_test

import (
	"fmt"
	"testing"
)

func Test_numSubarraysWithSum(t *testing.T) {
	goal := 2
	nums := []int{1, 0, 1, 0, 1}
	//goal := 0
	//nums := []int{0, 0, 0, 0, 0}
	fmt.Println(numSubarraysWithSum0(nums, goal))
}

func numSubarraysWithSum(nums []int, goal int) int {
	//思路使用两个左指针,在右指针不变的情况下,把两个左指针的的最大距离加入结果
	left1, left2 := 0, 0
	sum1, sum2 := 0, 0
	ans := 0
	for right, num := range nums {
		sum1 += num
		for left1 <= right && sum1 > goal {
			//右端点不变,满足要求的最左端
			sum1 -= nums[left1]
			left1++
		}
		sum2 += num
		for left2 <= right && sum2 >= goal {
			//右端点不变,满足要求的最右端
			sum2 -= nums[left2]
			left2++
		}
		//差值加入结果
		ans += left2 - left1
	}
	return ans
}
func numSubarraysWithSum0(nums []int, goal int) int {
	//前缀和+哈希
	//数组求和,结果递增,使用count记录,没个值出现的次数,出现0的时候上一个值会+1
	//sum-goal就是上一个满足的点,把满足点的个数加入结果
	n := len(nums)
	sum := 0
	count := map[int]int{}
	ans := 0
	for i := 0; i < n; i++ {
		count[sum] += 1
		sum += nums[i]
		//把满足点的个数加入结果
		ans += count[sum-goal]
	}
	return ans
}
