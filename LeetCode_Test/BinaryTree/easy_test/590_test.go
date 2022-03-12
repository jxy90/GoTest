package easy_test

import (
	"fmt"
	"testing"
)

func Test_590(t *testing.T) {
	root := &Node{
		Val: 1,
		Children: []*Node{{
			Val: 3,
			Children: []*Node{{
				Val: 5,
			}, {
				Val: 6,
			}},
		}, {
			Val: 2,
		}, {
			Val: 4,
		}},
	}
	fmt.Println(postorder(root))
	fmt.Println(postorder0(root))
}
func postorder(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node != nil {
			for _, child := range node.Children {
				dfs(child)
			}
			ans = append(ans, node.Val)
		}
	}
	dfs(root)
	return
}

func postorder0(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			dfs(child)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}
