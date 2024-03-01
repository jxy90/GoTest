package Array

import (
	"testing"
)

//常数时间删除-查找数组中的任意元素

//O(1) 时间插入、删除和获取随机元素
//https://leetcode.cn/problems/insert-delete-getrandom-o1/description/

//type RandomizedSet struct {
//	cache map[int]int
//	nums  []int
//}
//
//func Constructor1() RandomizedSet {
//	return RandomizedSet{
//		cache: make(map[int]int, 100),
//		nums:  make([]int, 0, 100),
//	}
//}
//
//func (this *RandomizedSet) Insert(val int) bool {
//	if _, ok := this.cache[val]; ok {
//		return false
//	}
//	this.cache[val] = len(this.nums)
//	this.nums = append(this.nums, val)
//	return true
//}
//
//func (this *RandomizedSet) Remove(val int) bool {
//	n := len(this.nums)
//	if index, ok := this.cache[val]; ok {
//		this.cache[this.nums[n-1]] = index
//		this.nums[index], this.nums[n-1] = this.nums[n-1], this.nums[index]
//		this.nums = this.nums[:n-1]
//		delete(this.cache, val)
//		return true
//	}
//	return false
//}
//
//func (this *RandomizedSet) GetRandom() int {
//	//fmt.Println(this.nums)
//	return this.nums[rand.Intn(len(this.nums))]
//}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

//黑名单中的随机数
//https://leetcode.cn/problems/random-pick-with-blacklist/description/
//type Solution struct {
//	nums []int
//}
//
//func Constructor(n int, blacklist []int) Solution {
//	nums := make([]int, 0, n)
//	cache := make(map[int]struct{})
//	for i := range blacklist {
//		cache[blacklist[i]] = struct{}{}
//	}
//	for i := 0; i < n; i++ {
//		if _, ok := cache[i]; !ok {
//			nums = append(nums, i)
//		}
//	}
//	return Solution{nums}
//}
//
//func (this *Solution) Pick() int {
//	return this.nums[rand.Intn(len(this.nums))]
//}

func Test_9(t *testing.T) {

}
