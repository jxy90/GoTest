package middle_test

import (
	"log"
	"testing"
)

func Test_429(t *testing.T) {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{3, []*Node{{5, nil}, {6, nil}}},
			{2, nil},
			{4, nil},
		},
	}
	log.Println(levelOrder(root))
}

func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	//当前层记录
	q := make([]*Node, 0)
	q = append(q, root)
	//下一层记录
	nextQ := make([]*Node, 0)
	//当前层index
	level := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		//当前层不存在,创建
		if len(ans) == level {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], cur.Val)
		for i := range cur.Children {
			nextQ = append(nextQ, cur.Children[i])
		}
		//当前层访问结束,进入下一层
		if len(q) == 0 {
			q = nextQ
			nextQ = []*Node{}
			level++
		}
	}
	return
}
