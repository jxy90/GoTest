package diffcult_test

import (
	"testing"
)

//1.对数组进行排序
func recoverTree(root *TreeNode) {
	list = []*TreeNode{}
	recoverTreeHelper(root)
	x, y := -1, -1
	for i := 0; i < len(list)-1; i++ {
		if list[i].Val > list[i+1].Val {
			y = i + 1
			if x == -1 {
				x = i
			}
		}
	}
	list[x].Val, list[y].Val = list[y].Val, list[x].Val
}

var list []*TreeNode

func recoverTreeHelper(root *TreeNode) {
	if root == nil {
		return
	}
	recoverTreeHelper(root.Left)
	list = append(list, root)
	recoverTreeHelper(root.Right)
}

//2.迭代中序
func recoverTree2(root *TreeNode) {
	var x, y, parent *TreeNode
	var stack []*TreeNode
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if parent != nil && parent.Val > current.Val {
			y = current
			if x == nil {
				x = parent
			} else {
				break
			}
		}
		parent = current
		current = current.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

//3. 递归中序
var x, y, parent *TreeNode

func recoverTree3(root *TreeNode) {
	x, y, parent = nil, nil, nil
	recoverTree3Findxy(root)
	x.Val, y.Val = y.Val, x.Val
}
func recoverTree3Findxy(root *TreeNode) {
	if root == nil {
		return
	}
	recoverTree3Findxy(root.Left)
	if parent != nil && parent.Val > root.Val {
		y = root
		if x == nil {
			x = parent
		} else {
			return
		}
	}
	parent = root
	recoverTree3Findxy(root.Right)
}

//4.Morris 中序遍历
func recoverTree4(root *TreeNode) {
	var x, y, pred, predecessor *TreeNode
	x, y, pred, predecessor = nil, nil, nil, nil

	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right

				if predecessor.Right == nil {
					predecessor.Right = root
					root = root.Left
				} else {
					// check for the swapped nodes
					if pred != nil && root.Val < pred.Val {
						y = root
						if x == nil {
							x = pred
						}
					}
					pred = root

					predecessor.Right = nil
					root = root.Right
				}
			}
		} else {
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root

			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func Test_recoverTree(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	recoverTree2(root)
}
