package hard_test

import (
	"container/heap"
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_trapRainWater(t *testing.T) {
	//fmt.Println(trapRainWater([][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}))
	fmt.Println(trapRainWater([][]int{{12, 13, 1, 12}, {13, 4, 13, 12}, {13, 8, 10, 12}, {12, 13, 12, 12}, {13, 13, 13, 13}}))
}

func trapRainWater(heightMap [][]int) (ans int) {
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	h := &hp{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				heap.Push(h, cell{i, j, heightMap[i][j]})
				//h.Push(cell{i, j, heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}
	dirctions := []int{-1, 0, 1, 0, -1}
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		//cur := h.Pop().(cell)
		for i := 0; i < 4; i++ {
			newX, newY := cur.x+dirctions[i], cur.y+dirctions[i+1]
			if 0 < newX && newX < m && 0 <= newY && newY < n && !visited[newX][newY] {
				if heightMap[newX][newY] < cur.h {
					ans += cur.h - heightMap[newX][newY]
				}
				visited[newX][newY] = true
				heap.Push(h, cell{newX, newY, CommonUtil.Max(cur.h, heightMap[newX][newY])})
				//h.Push(cell{newX, newY, CommonUtil.Max(cur.h, heightMap[newX][newY])})
			}
		}
	}
	return
}

type cell struct {
	x, y, h int
}

type hp []cell

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	return h[i].h < h[j].h
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(cell))
}
func (h *hp) Pop() interface{} {
	temp := *h
	v := temp[len(temp)-1]
	*h = temp[:len(temp)-1]
	return v
}
