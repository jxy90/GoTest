package hard_test

import "testing"

func Test_leastBricks(t *testing.T) {
	fmt.Println(leastBricks([][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}))
}

func leastBricks(wall [][]int) int {
	m := len(wall)
	cache := map[int]int{}
	for i := 0; i < m; i++ {
		temp := 0
		for j := 0; j < len(wall[i])-1; j++ {
			temp += wall[i][j]
			cache[temp]++
		}
	}
	maxCount := 0
	for _, v := range cache {
		maxCount = max(maxCount, v)
	}
	return m - maxCount
}

func leastBricks1(wall [][]int) int {
	m := len(wall)
	n := 0
	for _, v := range wall[0] {
		n += v
	}
	cache := map[int]map[int]bool{}
	for i := 0; i < m; i++ {
		cache[i] = map[int]bool{}
	}
	for i := 0; i < m; i++ {
		temp := 0
		for _, v := range wall[i] {
			temp += v
			column := cache[i]
			column[temp] = true
			//cache[i][temp] = temp
		}
	}
	ans := m
	for i := 1; i < n; i++ {
		count := 0
		for j := 0; j < m; j++ {
			if _, ok := cache[j][i]; !ok {
				count++
			}
		}
		ans = min(ans, count)
	}
	return ans
}
