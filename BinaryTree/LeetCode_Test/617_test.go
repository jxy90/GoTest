package leetcode_test

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	//if t1==nil {
	//	return t2
	//}
	//if t2==nil {
	//	return t1
	//}
	//sum :=t1.Val+t2.Val
	//t1.Val=sum
	//t1.Left= mergeTrees(t1.Left,t2.Left)
	//t1.Right=mergeTrees(t2.Left,t1.Left)
	//return t1

	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}

}

func mergeTreesHelper(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 != nil {
		return t2
	} else if t1 != nil && t2 == nil {
		return t1
	} else if t1 != nil && t2 != nil {
		return &TreeNode{
			Val: t1.Val + t2.Val,
		}
	}
	return nil
}
