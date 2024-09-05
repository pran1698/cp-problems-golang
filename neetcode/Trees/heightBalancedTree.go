package main

func checkBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	lD, lB := checkBalanced(root.Left)
	rD, rB := checkBalanced(root.Right)

	balanced := lB && rB
	if abs(lD-rD) > 1 {
		balanced = false
	}

	return 1 + max(lD, rD), balanced
}

func isBalanced(root *TreeNode) bool {
	_, isBalanced := checkBalanced(root)
	return isBalanced
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -1 * a
}
