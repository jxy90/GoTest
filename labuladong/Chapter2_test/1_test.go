package Chapter2_test

//数据结构的存储方式
//其实这只有两种
//1.链表（链式存储）
//2.数组（顺序存储）

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode) {
	// 前序遍历
	traverse(root.Left)
	// 中序遍历
	traverse(root.Right)
	// 后序遍历
}
