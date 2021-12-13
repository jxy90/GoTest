package easy_test

import (
	"fmt"
	"math"
	"testing"
)

func Test_0(t *testing.T) {
	//fmt.Println(countPoints("B0R0G0R9R0B0G0"))
	//fmt.Println(subArrayRanges([]int{1, 2, 3}))
	//fmt.Println(minimumRefill([]int{2,2,3,3},5,5))
	//fmt.Println(minimumRefill([]int{2,2,3,3},2,4))
	fmt.Println(maxTotalFruits([][]int{{2, 8}, {6, 3}, {8, 6}}, 5, 4))

}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	//f[5][4]=max(f[4][3],f[6,3])+pos[5]
	posCnt := map[int]int{}
	for _, fruit := range fruits {
		posCnt[fruit[0]] = fruit[1]
	}
	//f[5][4]=max(f[4][3],f[6,3])+pos[5]
	//f := make([][]int, startPos+k+1)
	//for i := range f {
	//	f[i] = make([]int, k+1)
	//}
	visited := map[int]bool{}
	ans := 0
	var dfs func(position int, step int, num int)
	dfs = func(position int, step int, num int) {
		if step == 0 {
			if num > ans {
				ans = num
				return
			}
		}
		cnt := 0
		if !visited[position] {
			cnt = posCnt[position]
		}
		if position+1 <= startPos+k {
			visited[position+1] = true
			dfs(position+1, step-1, num+cnt)
			visited[position+1] = false
		}
		if position-1 >= startPos-k {
			visited[position-1] = true
			dfs(position-1, step-1, num+cnt)
			visited[position-1] = true
		}
	}
	visited[startPos] = true
	dfs(startPos, k, 0)
	//for _, ints := range f {
	//	fmt.Println(ints)
	//}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	left, right := 0, len(plants)-1
	ans := 0
	nowA, nowB := capacityA, capacityB
	for left <= right {
		if left == right {
			if nowA >= nowB {
				if nowA < plants[left] {
					//nowA = capacityA
					ans++
					//nowA -= plants[left]
				}
			} else {
				if nowB < plants[left] {
					//nowA = capacityA
					ans++
					//nowA -= plants[left]
				}
			}
		} else {
			if nowA < plants[left] {
				nowA = capacityA
				ans++
			}
			nowA -= plants[left]
			if nowB < plants[right] {
				nowB = capacityB
				ans++
			}
			nowB -= plants[right]
		}
		left++
		right--
	}
	return ans
}

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	ans := int64(0)
	for i := 0; i < n; i++ {
		min, max := math.MaxInt64, math.MinInt64
		for j := i; j < n; j++ {
			if nums[j] > max {
				max = nums[j]
			}
			if nums[j] < min {
				min = nums[j]
			}
			ans += int64(max) - int64(min)
		}
	}
	return ans
}

func countPoints(rings string) int {
	m := map[byte]int{'R': 0, 'G': 1, 'B': 2}
	cache := map[byte]int{}
	n := len(rings) / 2
	ans := 0
	for i := 0; i < n; i++ {
		a, b := rings[i<<1], rings[i<<1+1]
		if cache[b] == 7 {
			continue
		}
		num := m[a]
		if (cache[b] | 1<<num) == 7 {
			ans++
		}
		cache[b] |= 1 << num
	}
	return ans
}
