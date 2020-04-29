package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	for k := range inorder {
		if inorder[k] == postorder[len(postorder)-1] {
			return &TreeNode{
				Val:   inorder[k],
				Left:  buildTree(inorder[:k], postorder[:k]),
				Right: buildTree(inorder[k+1:], postorder[k:len(postorder)-1]),
			}
		}
	}
	return nil
}
