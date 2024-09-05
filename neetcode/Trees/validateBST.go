package main

import "math"

func isValidBST(root *TreeNode) bool {

	return validate(root, math.MinInt, math.MaxInt)
}

func validate(node *TreeNode, lower int, upper int) bool {

	if node == nil {

		//empty node or empty tree is always valid
		return true
	}

	if (lower < node.Val) && (node.Val < upper) {

		// check if all tree nodes follow BST rule
		return validate(node.Left, lower, node.Val) && validate(node.Right, node.Val, upper)
	} else {

		// early reject when we find violation
		return false
	}

}
