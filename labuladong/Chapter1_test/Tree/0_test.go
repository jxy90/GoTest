package Tree

//二叉树（纲领篇）

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树的最大深度
//https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

//二叉树的直径
//https://leetcode.cn/problems/diameter-of-binary-tree/description/
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	diameterOfBinaryTreeHelper(root, &ans)
	return ans
}

func diameterOfBinaryTreeHelper(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := diameterOfBinaryTreeHelper(root.Left, ans)
	right := diameterOfBinaryTreeHelper(root.Right, ans)
	*ans = max(*ans, left+right)
	return max(left, right) + 1
}

//层序遍历
func levelTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
