package main

//分析：https://blog.csdn.net/qq_28114615/article/details/85715017
//O(n2)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root || q == root {
		return root
	}
	//if root == nil || p == root || q == root {
	//	return root
	//}

	Left := lowestCommonAncestor(root.Left, p, q)
	Right := lowestCommonAncestor(root.Right, p, q)

	if Left == nil && Right == nil {
		return nil
	} else if Left != nil && Right == nil {
		return Left
	} else if Left == nil && Right != nil {
		return Right
	}

	return root
}

//
//func findRootRoad(r, p *TreeNode, road []*TreeNode, result *[]*TreeNode) {
//	if r == nil {
//		return
//	}
//
//	road = append(road, r)
//
//	if r == p {
//		*result = make([]*TreeNode, len(road))
//		copy(*result, road)
//		return
//	}
//
//	if r.Left != nil {
//		findRootRoad(r.Left, p, road, result)
//	}
//	if r.Right != nil {
//		findRootRoad(r.Right, p, road, result)
//	}
//}
//
//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	pR := make([]*TreeNode, 0)
//	qR := make([]*TreeNode, 0)
//	findRootRoad(root, p, make([]*TreeNode, 0), &pR)
//	findRootRoad(root, q, make([]*TreeNode, 0), &qR)
//	var result *TreeNode
//	for i := 0; i < len(pR) && i < len(qR); i++ {
//		if pR[i] == qR[i] {
//			result = pR[i]
//		}
//	}
//	return result
//}
