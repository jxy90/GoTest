package middle_test

import (
	"fmt"
	"testing"
)

func Test_getDistances(t *testing.T) {
	fmt.Println(getDistances([]int{2, 1, 3, 1, 2, 3, 3}))
}

func getDistances(arr []int) []int64 {
	n := len(arr)
	cache := map[int][]int{}
	ans := make([]int64, n)
	for i, v := range arr {
		if v == 1 {
			fmt.Println()
		}
		for _, index := range cache[v] {
			ans[index] += int64(i - index)
			ans[i] += int64(i - index)
		}
		cache[v] = append(cache[v], i)
	}
	return ans
}

func getDistances0(arr []int) []int64 {
	cache := map[int][]int{}
	for i, v := range arr {
		cache[v] = append(cache[v], i) // 记录相同元素的位置
	}
	ans := make([]int64, len(arr))
	for _, p := range cache {
		sum := 0
		for _, i := range p {
			sum += i - p[0]
		}
		ans[p[0]] = int64(sum) // 最左侧元素的间隔和
		for i, n := 1, len(p); i < n; i++ {
			sum += (2*i - n) * (p[i] - p[i-1]) // 计算下一个相同元素的间隔和（考虑变化量）
			ans[p[i]] = int64(sum)
		}
	}
	return ans
}
