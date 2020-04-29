package main

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	for k := range inorder {
		if inorder[k] == preorder[0] {
			return &TreeNode{
				Val:   inorder[k],
				Left:  buildTree2(preorder[1:k+1], inorder[:k]),
				Right: buildTree2(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}
