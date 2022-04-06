package middle_test

import (
	"sort"
	"testing"
)

func Test_954(t *testing.T) {
}

func canReorderDoubled(arr []int) bool {
	//缓存数字出现的次数
	cnt := make(map[int]int, len(arr))
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]%2 == 1 {
		return false
	}
	//如果为true的话，len(cnt)==len(arr)/2
	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })
	//不存在的时候，说明不符合
	for _, x := range vals {
		if cnt[2*x] < cnt[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		cnt[2*x] -= cnt[x]
	}
	return true
}
