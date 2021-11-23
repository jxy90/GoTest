package middle_test

import "math/rand"

type Solution384 struct {
	original []int
	nums     []int
}

func Constructor384(nums []int) Solution384 {
	return Solution384{
		original: append([]int{}, nums...),
		nums:     nums,
	}
}

func (this *Solution384) Reset() []int {
	copy(this.nums, this.original)
	return this.original
}

func (this *Solution384) Shuffle() []int {
	//temp := make([]int, len(this.nums))
	//for i := range temp {
	//	j := rand.Intn(len(this.nums))
	//	temp[i] = this.nums[j]
	//	this.nums = append(this.nums[:j], this.nums[j+1:]...)
	//}
	//this.nums = temp
	n := len(this.nums)
	for i := range this.nums {
		j := i + rand.Intn(n-i)
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	}
	return this.nums
}

//Your Solution object will be instantiated and called as such:
