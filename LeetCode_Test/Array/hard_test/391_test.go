package hard_test

import (
	"fmt"
	"math"
	"testing"
)

func Test_isRectangleCover(t *testing.T) {
	fmt.Println(isRectangleCover([][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}))
}

func isRectangleCover(rectangles [][]int) bool {
	m := []int{math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32}
	points := make(map[Point]int)
	sumArea := 0
	for _, rectangle := range rectangles {
		x, y, a, b := rectangle[0], rectangle[1], rectangle[2], rectangle[3]
		//记录point
		points[Point{x, y}]++
		points[Point{x, b}]++
		points[Point{a, b}]++
		points[Point{a, y}]++
		//面积和
		area := (a - x) * (b - y)
		sumArea += area
		//最大的矩形
		m[0] = min(m[0], x)
		m[1] = min(m[1], y)
		m[2] = max(m[2], a)
		m[3] = max(m[3], b)
	}
	//面积不相等,返回false
	if sumArea != (m[2]-m[0])*(m[3]-m[1]) {
		return false
	}
	//移除四个角
	points[Point{m[0], m[1]}]++
	points[Point{m[0], m[3]}]++
	points[Point{m[2], m[1]}]++
	points[Point{m[2], m[3]}]++
	for _, v := range points {
		if v != 2 && v != 4 {
			return false
		}
	}
	return true
}

type Point struct {
	x, y int
}
