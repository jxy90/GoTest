package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_isNStraightHand(t *testing.T) {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 4))
}

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	//不能整除
	if n%groupSize > 0 {
		return false
	}
	sort.Ints(hand)
	cache := map[int]int{}
	//记录每个数字的次数
	for _, v := range hand {
		cache[v]++
	}
	for _, k := range hand {
		//当前数字用完了，下一个
		if cache[k] == 0 {
			continue
		}
		//对连续的groupSize个数字进行-1操作，不能减的返回false
		for i := k; i < k+groupSize; i++ {
			if cache[i] == 0 {
				return false
			}
			cache[i]--
		}
	}
	return true
}
