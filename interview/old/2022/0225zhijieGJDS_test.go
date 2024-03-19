package _022

import (
	"fmt"
	"testing"
)

//redis为什么是用SDS，Harsh的具体操作和新数据进入怎么处理？
//mysql优化器如何选择索引
//es有数据量限制么
//https://www.cnblogs.com/25-lH/p/10973309.html
//mysql 跳过limit 优化
//算法:数字全排列
//https://www.cnblogs.com/zhangbiao97/p/13062036.html

func Test_gjds(t *testing.T) {
	//fmt.Println("123")
	fmt.Println(findNumber([]int{1, 2, 3, 4, 5}))
	fmt.Println(findNumber([]int{1, 2, 3, 5, 4}))
	fmt.Println(findNumber([]int{1, 2, 4, 3, 5}))
	fmt.Println(findNumber([]int{1, 2, 4, 5, 3}))
	fmt.Println(findNumber([]int{1, 2, 5, 3, 4}))
	fmt.Println(findNumber([]int{1, 5, 4, 3, 2}))
}

func findNumber(nums []int) []int {
	//1.从后向前找到逆序和顺序交界的index
	index := getIndex(nums)
	if index == 0 {
		return nil
	}
	//2.把逆序前一位的数字,和逆序位的中刚大于他的数字交换
	var copyNums = append([]int{}, nums...)
	exchangeNum(copyNums, index)
	//3.把交换后的逆序换位顺序
	reverse(copyNums, index)
	return copyNums
}

func reverse(nums []int, index int) {
	i, j := index, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func exchangeNum(nums []int, index int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[index-1] {
			nums[i], nums[index-1] = nums[index-1], nums[i]
			break
		}
	}
}

func getIndex(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			return i
		}
	}
	return 0
}
