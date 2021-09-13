package middle_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_numberOfBoomerangs(t *testing.T) {
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}))
	//fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
}

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for _, p := range points {
		cache := map[int]int{}
		for _, q := range points {
			distance := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cache[distance]++
		}
		for _, c := range cache {
			ans += c * (c - 1)
		}
	}
	return ans
}

//ETL
func numberOfBoomerangs0(points [][]int) int {
	n := len(points)
	temp := make([][]int, 0)
	ans := 0
	visited := map[int]bool{}
	var dfs func()
	dfs = func() {
		if len(temp) == 3 {
			ans++
			fmt.Println(temp)
			return
		}
		for i := 0; i < n; i++ {
			if !visited[i] {
				if len(temp) < 2 || check(temp, points[i]) {
					visited[i] = true
					temp = append(temp, points[i])
					dfs()
					temp = temp[:len(temp)-1]
					visited[i] = false
				}
			}
		}
	}
	dfs()
	return ans
}

func check(temp [][]int, z []int) bool {
	x, y := temp[0], temp[1]
	a, b, c, d := CommonUtil.Abs(x[0]-y[0]), CommonUtil.Abs(x[1]-y[1]), CommonUtil.Abs(x[0]-z[0]), CommonUtil.Abs(x[1]-z[1])
	return a*a+b*b == c*c+d*d
}
