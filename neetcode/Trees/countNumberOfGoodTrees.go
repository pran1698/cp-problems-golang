package main

func countGoodNodes(root *TreeNode, maxV int) int {
	if root == nil {
		return 0
	}

	if maxV <= root.Val {
		return 1 + countGoodNodes(root.Left, root.Val) + countGoodNodes(root.Right, root.Val)
	}

	return countGoodNodes(root.Left, maxV) + countGoodNodes(root.Right, maxV)
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countGoodNodes(root, -100000)
}
