package hard_test

import (
	"testing"
)

func Test_trapRainWater(t *testing.T) {
	//fmt.Println(trapRainWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	waterMap := make([][]int, m)
	for i := range waterMap {
		waterMap[i] = make([]int, n)
		for j := range waterMap[i] {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				waterMap[i][j] = heightMap[i][j]
			}
		}
	}
	for i := 1; i < m-2; i++ {
		for j := 1; j < n-2; j++ {
			waterMap[i][j] = 1
		}
	}
	return ans
}
