package middle_test

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queueT := []*TreeNode{root}
	queueI := []int{1}
	for i := 0; i < len(queueT); i++ {

		node := queueT[i]
		nodeIndex := queueI[i]
		if node.Left != nil {
			queueT = append(queueT, node.Left)
			queueI = append(queueI, 2*nodeIndex)
		}
		if node.Right != nil {
			queueT = append(queueT, node.Right)
			queueI = append(queueI, 2*nodeIndex+1)
		}
	}
	return len(queueI) == queueI[len(queueI)-1]
}
