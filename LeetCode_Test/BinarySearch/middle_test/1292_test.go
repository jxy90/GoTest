package middle_test

import "github.com/jxy90/GoTest/Utils/CommonUtil"

//func Test_shipWithinDays(t *testing.T) {
//	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
//}

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])
	sum := make([][]int, n+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i+1][j+1] = mat[i][j] + sum[i+1][j] + sum[i][j+1] - sum[i][j]
		}
	}
	maxlength := CommonUtil.Min(m, n)
	l, r := 0, maxlength
	for l < r {
		mid := l + (r-l+1)>>1
		if maxSideLengthHelper(sum, mid, threshold) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func maxSideLengthHelper(sum [][]int, mid, threshold int) bool {
	m := len(sum)
	n := len(sum[0])
	for i := 1; i+mid-1 < m; i++ {
		for j := 1; j+mid-1 < n; j++ {
			diff := sum[i+mid-1][j+mid-1] - sum[i+mid-1][j-1] - sum[i-1][j+mid-1] + sum[i-1][j-1]
			if diff <= threshold {
				return true
			}
		}
	}
	return false
}
