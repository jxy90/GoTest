package BackPack

import (
	"fmt"
	"math"
	"testing"
)

//787. K 站中转内最便宜的航班
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	//memo := make(map[int]int)
	temp := make(map[int][][]int)
	for _, flight := range flights {
		if _, ok := temp[flight[0]]; !ok {
			temp[flight[0]] = make([][]int, 0)
		}
		temp[flight[0]] = append(temp[flight[0]], flight)
	}
	res := math.MaxInt32
	var helper func(src, k, price int)
	helper = func(src, k, price int) {
		if src == dst {
			res = min(res, price)
			return
		}
		if k < 0 {
			return
		}
		for _, flight := range temp[src] {
			price += flight[2]
			//if val, ok := memo[flight[1]]; ok && val <= price {
			//	continue
			//}
			//memo[flight[1]] = price
			helper(flight[1], k-1, price)
			price -= flight[2]
		}
	}
	helper(src, k, 0)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func Test_3(t *testing.T) {
	//fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1))
}
