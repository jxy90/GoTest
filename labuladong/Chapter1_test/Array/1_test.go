package Array

//小而美的算法技巧：前缀和数组

//一维数组中的前缀和
//https: //leetcode.cn/problems/range-sum-query-immutable/

//type NumArray struct {
//	sum []int
//}
//
//func Constructor(nums []int) NumArray {
//	sum := make([]int, len(nums)+1)
//	sum[1] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		sum[i+1] = nums[i] + sum[i]
//	}
//	return NumArray{sum: sum}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	return this.sum[right+1] - this.sum[left]
//}

//二维矩阵中的前缀和
//https://leetcode.cn/problems/range-sum-query-2d-immutable/

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	//构建前缀和
	sum := make([][]int, len(matrix)+1)
	for i := range sum {
		sum[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i < len(sum); i++ {
		for j := 1; j < len(sum[i]); j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{sum: sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row1][col2+1] - this.sum[row2+1][col1] + this.sum[row1][col1]
}
