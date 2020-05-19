package diffcult_test

func minCameraCover(root *TreeNode) int {

	if root == nil {
		return 0
	}
	left := minCameraCoverDFS(root.Left)
	right := minCameraCoverDFS(root.Right)
	return root.Val + left + right - 1
}

func minCameraCoverDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minCameraCoverDFS(root.Left)
	right := minCameraCoverDFS(root.Right)
	return root.Val + left + right - 1
}
