package Chapter2_test

import (
	"fmt"
	"testing"
)

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

func Test1(t *testing.T) {
	call()
}

type Test struct {
	Max int
}

func (t *Test) Println() {
	fmt.Println(t.Max)
}

func deferExec(f func()) {
	f()
}

func call() {
	var t *Test
	defer deferExec(t.Println)

	t = new(Test)
}
