package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLeft := dfs(root.Left)
	maxRight := dfs(root.Right)
	maxLeft = max(maxLeft, 0)
	maxRight = max(maxRight, 0)

	// compute max path sum WITH spilt
	res = max(res, root.Val+maxLeft+maxRight)

	// return value is value of path without split
	return root.Val + max(maxLeft, maxRight)
}

func maxPathSum(root *TreeNode) int {
	res = -1000000

	dfs(root)
	return res
}
