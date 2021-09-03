package middle_test

import (
	"testing"
)

func Test_canEat(t *testing.T) {
	fmt.Println(canEat([]int{46, 5, 47, 48, 43, 34, 15, 26, 11, 25, 41, 47, 15, 25, 16, 50, 32, 42, 32, 21, 36, 34, 50, 45, 46, 15, 46, 38, 50, 12, 3, 26, 26, 16, 23, 1, 4, 48, 47, 32, 47, 16, 33, 23, 38, 2, 19, 50, 6, 19, 29, 3, 27, 12, 6, 22, 33, 28, 7, 10, 12, 8, 13, 24, 21, 38, 43, 26, 35, 18, 34, 3, 14, 48, 50, 34, 38, 4, 50, 26, 5, 35, 11, 2, 35, 9, 11, 31, 36, 20, 21, 37, 18, 34, 34, 10, 21, 8, 5},
		[][]int{{85, 54, 42}}))
	fmt.Println(canEat([]int{7, 11, 5, 3, 8}, [][]int{{2, 2, 6}, {4, 2, 4}, {2, 13, 1000000000}}))
	fmt.Println(canEat([]int{7, 4, 5, 3, 8}, [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 1000000000}}))
	fmt.Println(canEat([]int{5, 2, 6, 4, 1}, [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}}))
}

func canEat(candiesCount []int, queries [][]int) []bool {
	sum := make([]int, len(candiesCount)+1)
	for i := 1; i <= len(candiesCount); i++ {
		sum[i] = sum[i-1] + candiesCount[i-1]
	}
	n := len(queries)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		fType, fDay, fCap := queries[i][0], queries[i][1]+1, queries[i][2]
		minDay := sum[fType]/fCap + 1
		maxDay := sum[fType+1]
		if minDay <= fDay && fDay <= maxDay {
			ans[i] = true
		}
	}
	return ans
}

func canEat1(candiesCount []int, queries [][]int) []bool {
	cacheBeginDay := make(map[int]int)
	cacheEnd := make(map[int]int)
	index := 0
	for i, value := range candiesCount {
		cacheBeginDay[i] = index
		index += value - 1
		cacheEnd[i] = index
		index++
	}
	n := len(queries)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = canEatValidate(queries[i], cacheBeginDay, cacheEnd)
	}
	return ans
}

func canEatValidate(query []int, cacheBegin, cacheEnd map[int]int) bool {
	left := query[1]
	right := query[1] * query[2]
	begin := cacheBegin[query[0]]
	end := cacheEnd[query[0]]
	if (left <= begin && begin <= right) || (left <= end && end <= right) || (begin <= left && right <= end) {
		return true
	}
	return false
}

func canEat0(candiesCount []int, queries [][]int) []bool {
	cache := make(map[int]int)
	index := 0
	for i, value := range candiesCount {
		for j := 0; j < value; j++ {
			cache[index] = i
			index++
		}
	}
	m := len(cache)
	n := len(queries)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		left := queries[i][1]
		right := (queries[i][1] + 1) * queries[i][2]
		if 0 <= left && left < m && cache[left] <= queries[i][0] && (right >= m || (0 <= right && right < m && cache[right] >= queries[i][0])) {
			ans[i] = true
		}
	}
	return ans
}
