package LeetCode_Test

import "testing"

func Test_getSkyline(t *testing.T) {
	fmt.Println(getSkyline([][]int{{0, 2, 3}, {2, 5, 3}}))
	//fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
}

//func getSkyline(buildings [][]int) [][]int {
//
//}

func getSkyline(buildings [][]int) [][]int {
	max := buildings[0][1]
	for _, building := range buildings {
		if max < building[1] {
			max = building[1]
		}
	}
	height := make([]int, max+2)
	for _, building := range buildings {
		start, end, h := building[0], building[1], building[2]
		for i := start; i < end; i++ {
			if height[i] < h {
				height[i] = h
			}
		}
	}
	ans := make([][]int, 0)
	for i := 0; i < len(height)-1; i++ {
		if i == 0 && height[i] > 0 {
			ans = append(ans, []int{i, height[i]})
		}
		if height[i] != height[i+1] {
			ans = append(ans, []int{i + 1, height[i+1]})
		}
	}
	return ans
}
