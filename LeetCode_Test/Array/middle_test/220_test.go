package middle_test

import (
	"testing"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	mp := map[int]int{}
	for i, num := range nums {
		id := getID(num, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if v, has := mp[id-1]; has && abs(v-num) <= t {
			return true
		}
		if v, has := mp[id+1]; has && abs(v-num) <= t {
			return true
		}
		mp[id] = num
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsNearbyAlmostDuplicate2(nums []int, k, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

func Test_containsNearbyAlmostDuplicate(t *testing.T) {

}
