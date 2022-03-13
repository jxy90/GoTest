package old

import (
	"fmt"
	"sort"
	"testing"
)

func Test_28(t *testing.T) {
	//fmt.Println(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1))
	//fmt.Println(findKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 1))

	//fmt.Println(digArtifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}}))
	//fmt.Println(digArtifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}, {1, 1}}))

	//fmt.Println(maximumTop([]int{5, 2, 2, 4, 0, 6}, 4))
	//fmt.Println(maximumTop([]int{2}, 1))
	//fmt.Println(maximumTop([]int{68, 76, 53, 73, 85, 87, 58, 24, 48, 59, 38, 80, 38, 65, 90, 38, 45, 22, 3, 28, 11}, 1))
	//fmt.Println(maximumTop([]int{18}, 3))

	//fmt.Println(minimumWeight(3, [][]int{{0, 1, 1}, {2, 1, 1}}, 0, 1, 2))
	fmt.Println(minimumWeight(6, [][]int{{0, 2, 2}, {0, 5, 6}, {1, 0, 3}, {1, 4, 5}, {2, 1, 1}, {2, 3, 3}, {2, 3, 4}, {3, 4, 2}, {4, 5, 1}}, 0, 1, 5))
}

//连续的自然数,移除两个,乱序
//map min,max.

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int {
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	for _, edge := range edges {
		x, y, w := edge[0], edge[1], edge[2]
		f[x][y] = w
	}

	var helperj func(j int)
	var helperi func(j int)
	helperj = func(j int) {
		for i := 0; i < n; i++ {
			for k := 0; k < n; k++ {
				if f[i][k] != -1 && f[k][j] != -1 && f[i][j] > f[i][k]+f[k][j] {
					f[i][j] = f[i][k] + f[k][j]
					helperi(i)
					helperj(j)
				}
			}
		}
	}
	helperi = func(i int) {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if f[i][k] != -1 && f[k][j] != -1 && f[i][j] > f[i][k]+f[k][j] {
					f[i][j] = f[i][k] + f[k][j]
					helperi(i)
					helperj(j)
				}
			}
		}
	}
	//s1->s2-dest
	for j := 0; j < n; j++ {
		helperi(j)
		helperj(j)
	}
	if f[src1][src2] == -1 || f[src2][dest] == -1 {
		return -1
	}
	return f[src1][src2] + f[src2][dest]
}

func maximumTop(nums []int, k int) int {
	if k == 0 {
		return nums[0]
	}
	if k == 1 {
		if len(nums) <= 1 {
			return -1
		} else {
			return nums[1]
		}
	}
	if len(nums) == 1 && k%2 == 1 {
		return -1
	}
	deleteMax := -1
	index := 0
	for k > 1 && index < len(nums) {
		if nums[index] > deleteMax {
			deleteMax = nums[index]
		}
		index++
		k--
	}
	if index == len(nums) {
		return deleteMax
	}
	if index == 0 {
		if len(nums) == 1 {
			return -1
		} else {
			return nums[index]
		}
	} else if index == len(nums)-1 {
		return deleteMax
	}
	if deleteMax > nums[index+1] {
		return deleteMax
	} else {
		return nums[index+1]
	}
}

func digArtifacts(n int, artifacts [][]int, dig [][]int) (ans int) {
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
	}
	for _, ints := range dig {
		x, y := ints[0], ints[1]
		d[x][y] = 1
	}
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = d[i-1][j-1] + f[i][j-1] + f[i-1][j] - f[i-1][j-1]
		}
	}
	for _, artifact := range artifacts {
		x1, y1, x2, y2 := artifact[0], artifact[1], artifact[2]+1, artifact[3]+1
		size := (x2 - x1) * (y2 - y1)
		a := f[x2][y2] - f[x2][y1] - f[x1][y2] + f[x1][y1]
		if a == size {
			ans++
		}

	}
	return ans
}

func findKDistantIndices(nums []int, key int, k int) (ans []int) {
	temp := make([]int, 0)
	for i, num := range nums {
		if num == key {
			temp = append(temp, i)
		}
	}
	t := make([]int, 0)
	for _, v := range temp {
		for i := v - k; i <= v+k; i++ {
			if i >= 0 && i < len(nums) {
				t = append(t, i)
			}
		}
	}
	sort.Ints(t)
	for i := range t {
		if i == 0 || t[i] != t[i-1] {
			ans = append(ans, t[i])
		}
	}
	return
}
