package middle_test

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := pathSumHelper(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}
func pathSumHelper(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return 0
	}
	if targetSum == root.Val {
		res++
	}
	res += pathSumHelper(root.Left, targetSum-root.Val)
	res += pathSumHelper(root.Right, targetSum-root.Val)
	return res
}
