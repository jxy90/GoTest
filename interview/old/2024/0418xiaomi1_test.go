package _024

import (
	"testing"
)

func Test_(t *testing.T) {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	visited := make(map[int]bool)
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil || right == nil {
		return nil
	} else if left != nil || right == nil {
		return right
	} else if left == nil || right != nil {
		return left
	}
	return root
}
