package Struct

import "math/rand"

type Solution struct {
	nums []int
	m    map[int][]int
}

func ConstructorPick(nums []int) Solution {
	cache := make(map[int][]int)
	for i, num := range nums {
		if _, ok := cache[num]; !ok {
			cache[num] = []int{i}
		} else {
			cache[num] = append(cache[num], i)
		}
	}
	return Solution{nums: nums, m: cache}
}

func (this *Solution) Pick(target int) int {
	//nums := this.m[target]
	//index := rand.Intn(len(nums))
	//return nums[index]
	cnt := 0
	ans := 0
	for i, num := range this.nums {
		if num == target {
			cnt++ // 第 cnt 次遇到 target
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return ans
}
