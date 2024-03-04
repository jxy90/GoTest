package Tree

//翻转二叉树
//https://leetcode.cn/problems/invert-binary-tree/description/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

//充每个节点的下一个右侧节点指针
//https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			cur := q[0]
			q = q[1:]
			if i < length-1 {
				cur.Next = q[0]
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return root
}

//二叉树展开为链表
//https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
