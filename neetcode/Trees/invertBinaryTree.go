package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	t1 := invertTree(root.Left)
	t2 := invertTree(root.Right)

	root.Right = t1
	root.Left = t2
	return root
}
