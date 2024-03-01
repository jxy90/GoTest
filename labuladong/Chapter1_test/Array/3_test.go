package Array

import (
	"fmt"
	"testing"
)

//二维数组的花式遍历技巧

//反转字符串中的单词
func reverseWords(s string) string {
	b := []byte(s)
	//去除头空格,让fast指向第一个不为空的位置
	slow, fast := 0, 0
	for fast < len(b) && b[fast] == ' ' {
		fast++
	}
	//去除中间多余空格
	for fast < len(b) {
		if fast-1 > 0 && fast < len(b) && b[fast] == b[fast-1] && b[fast] == ' ' {
			fast++
			continue
		}
		b[slow] = b[fast]
		slow++
		fast++
	}
	//去除后续不需要的子串
	if slow-1 > 0 && b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}
	//b = b[:slow]

	//翻转整个b
	reverse(&b, 0, len(b)-1)

	//安装单词分别翻转
	for i := 0; i < len(b); i++ {
		j := i
		for j < len(b) && b[j] != ' ' {
			j++
		}
		reverse(&b, i, j-1)
		i = j
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

//旋转图像
//https://leetcode.cn/problems/rotate-image/description/
func rotate(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := range matrix {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}

//螺旋矩阵
//https://leetcode.cn/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	up, down := 0, m-1
	left, right := 0, n-1
	ans := make([]int, 0, m*n)
	for len(ans) < m*n {
		//右
		if up <= down {
			for i := left; i <= right; i++ {
				ans = append(ans, matrix[up][i])
			}
			up++
		}
		//下
		if left <= right {
			for i := up; i <= down; i++ {
				ans = append(ans, matrix[i][right])
			}
			right--
		}

		//左
		if up <= down {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[down][i])
			}
			down--
		}

		//上
		if left <= right {
			for i := down; i >= up; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}
	return ans
}

//螺旋矩阵 II
//https://leetcode.cn/problems/spiral-matrix-ii/description/
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	up, down := 0, n-1
	left, right := 0, n-1
	for i := 1; i <= n*n; {
		if up <= down {
			for j := left; j <= right; j++ {
				ans[up][j] = i
				i++
			}
			up++
		}
		if left <= right {
			for j := up; j <= down; j++ {
				ans[j][right] = i
				i++
			}
			right--
		}
		if up <= down {
			for j := right; j >= left; j-- {
				ans[down][j] = i
				i++
			}
			down--
		}
		if left <= right {
			for j := down; j >= up; j-- {
				ans[j][left] = i
				i++
			}
			left++
		}
	}
	return ans
}

func Test_3(t *testing.T) {
	//fmt.Println(reverseWords("  hello world  "))
	fmt.Println(spiralArray([][]int{}))
}

func spiralArray(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	m, n := len(array), len(array[0])
	ans := make([]int, 0, m*n)
	up, down := 0, m-1
	left, right := 0, n-1
	for len(ans) < m*n {
		if up <= down {
			for i := left; i <= right; i++ {
				ans = append(ans, array[up][i])
			}
			up++
		}
		if left <= right {
			for i := up; i <= down; i++ {
				ans = append(ans, array[i][right])
			}
			right--
		}
		if up <= down {
			for i := right; i >= left; i-- {
				ans = append(ans, array[down][i])
			}
			down--
		}
		if left <= right {
			for i := down; i >= up; i-- {
				ans = append(ans, array[i][left])
			}
			left++
		}
	}
	return ans
}
