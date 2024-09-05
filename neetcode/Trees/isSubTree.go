package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if isSameTree(root, subRoot) {
		return true
	} else if root == nil && subRoot != nil {
		return false
	}

	if isSubtree(root.Left, subRoot) {
		return true
	}

	if isSubtree(root.Right, subRoot) {
		return true
	}

	return false
}
